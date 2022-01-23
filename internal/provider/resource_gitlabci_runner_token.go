package provider

// https://docs.gitlab.com/ce/api/runners.html#register-a-new-runner

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/parnurzeal/gorequest"
)

const resourceGitlabRunnerDescription = "The `gitlabci_runner` resource " +
	`allows the trivial creation of a runner
token using a runner registration token, without requiring authentication to
the GitLab instance itself.  It does this by using the [runner registration
token](https://docs.gitlab.com/runner/register/) of the
project/group/instance, rather than the authentication credentials of any
specific user.

Note that, once created, the registration is immutable: any changes will
result in the resource being destroyed and recreated.

See the [Registering a runner](https://registry.terraform.io/providers/rsrchboy/gitlabci/latest/docs/guides/registering-a-runner)
guide for more information.`

func resourceGitlabRunner() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceGitlabRunnerCreate,
		ReadContext:   resourceGitlabRunnerRead,
		DeleteContext: resourceGitlabRunnerDelete,
		Description:   resourceGitlabRunnerDescription,

		Schema: map[string]*schema.Schema{
			"registration_token": {
				Type:        schema.TypeString,
				ForceNew:    true,
				Required:    true,
				Sensitive:   true,
				StateFunc:   hashSum,
				Description: "Runner registration token (shared, group, or project)",
			},
			"token": {
				Type:        schema.TypeString,
				Computed:    true,
				Sensitive:   true,
				Description: "Generated (registered) runner token",
			},
			"runner_id": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Runner ID",
			},
			"access_level": {
				Type:         schema.TypeString,
				ForceNew:     true,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"not_protected", "ref_protected"}, true),
				Description:  "Run against all refs, or protected only",
			},
			"locked": {
				Type:        schema.TypeBool,
				ForceNew:    true,
				Optional:    true,
				Default:     true,
				Description: "Lock runner to project",
			},
			"maximum_timeout": {
				Type:         schema.TypeInt,
				ForceNew:     true,
				Optional:     true,
				ValidateFunc: validation.IntAtLeast(10 * 60),
				Description:  "Maximum timeout for jobs",
			},
			"tags": {
				Type:        schema.TypeSet,
				Optional:    true,
				ForceNew:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Set:         schema.HashString,
				Description: "List of tags for the runner",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Default:     "",
				Description: "Runner description",
			},
			"active": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Default:     true,
				Description: "Create the runner active, or paused?",
			},
			"run_untagged": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Default:     true,
				Description: "Take and run untagged jobs?",
			},
		},
	}
}

func resourceGitlabRunnerCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	baseURL := meta.(string)

	tflog.Trace(ctx, "create gitlab runner token")

	type RegisterOptions struct {
		Token          string   `json:"token"`
		Description    string   `json:"description,omitempty"`
		Active         bool     `json:"active"`
		Locked         bool     `json:"locked"`
		RunUntagged    bool     `json:"run_untagged"`
		TagList        []string `json:"tag_list,omitempty"`
		AccessLevel    string   `json:"access_level,omitempty"`
		MaximumTimeout int      `json:"maximum_timeout,omitempty"`
		// info hash
	}

	type registrationResponse struct {
		ID    int    `json:"id"`
		Token string `json:"token"`
	}

	registrationToken := d.Get("registration_token").(string)

	query := RegisterOptions{
		Token:          registrationToken,
		Description:    d.Get("description").(string),
		RunUntagged:    d.Get("run_untagged").(bool),
		Active:         d.Get("active").(bool),
		Locked:         d.Get("locked").(bool),
		MaximumTimeout: d.Get("maximum_timeout").(int),
	}

	if v, ok := d.GetOk("tags"); ok {
		query.TagList = *(stringSetToStringSlice(v.(*schema.Set)))
	}

	url := baseURL + "/runners"

	j, _ := json.Marshal(query)
	tflog.Trace(ctx, "create gitlab runner query: %s", j)

	req := gorequest.
		New().
		Post(url).
		Send(query)

	// TODO other registration options...

	var runnerDetails registrationResponse
	resp, _, errs := req.EndStruct(&runnerDetails)

	for _, err := range errs {
		// FIXME
		return diag.FromErr(err)
	}

	if resp.StatusCode != 201 {
		return diag.Errorf("bad response (%d): %s", resp.StatusCode, resp.Status)
	}

	d.SetId(fmt.Sprintf("%d", runnerDetails.ID))
	d.Set("token", runnerDetails.Token)
	d.Set("runner_id", runnerDetails.ID)

	return nil
}

func resourceGitlabRunnerRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	baseURL := meta.(string)

	runnerID, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	// FIXME probably ought to VerifyRegisteredRunner() here first

	tflog.Trace(ctx, "read gitlab runner %d", runnerID)

	url := baseURL + "/runners/verify"
	req := gorequest.
		New().
		Post(url).
		Send("token=" + d.Get("token").(string))
	resp, _, errs := req.End()

	for _, err := range errs {
		// FIXME
		return diag.FromErr(err)
	}

	if resp.StatusCode == 200 {
		// all good!
		return nil
	}

	return diag.Errorf("bad response (%d): %s", resp.StatusCode, resp.Status)
}

func resourceGitlabRunnerDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	baseURL := meta.(string)

	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	tflog.Trace(ctx, "Delete gitlab runner %d", id)

	url := baseURL + "/runners"
	req := gorequest.
		New().
		Delete(url).
		Send("token=" + d.Get("token").(string))
	resp, _, errs := req.End()

	for _, err := range errs {
		// FIXME
		return diag.FromErr(err)
	}

	if resp.StatusCode == 204 {
		// all good!
		return diag.FromErr(err)
	}

	return diag.Errorf("bad response (%d): %s", resp.StatusCode, resp.Status)
}
