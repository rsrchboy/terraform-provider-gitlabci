package provider

// https://docs.gitlab.com/ce/api/runners.html#register-a-new-runner

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	// we're embedding stuff, and the linter _really_ wants us to justify it
	_ "embed"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"gitlab.com/rsrchboy/terraform-provider-gitlabci/third_party/gitlab/runner/config"
)

//go:embed resource_runner_token.md
var resourceGitlabRunnerDescription string

const fnDesc = " *Changing this attribute forces the recreation of the resource.*"

func resourceGitlabRunner() *schema.Resource {
	s := &schema.Resource{
		CreateContext: resourceGitlabRunnerCreate,
		ReadContext:   resourceGitlabRunnerRead,
		DeleteContext: resourceGitlabRunnerDelete,
		UpdateContext: resourceGitlabRunnerUpdate,
		Description:   resourceGitlabRunnerDescription,

		Schema: map[string]*schema.Schema{
			"registration_token": {
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				StateFunc:   hashSum,
				Description: "Runner registration token; see [Registering Runners](https://docs.gitlab.com/runner/register) for more information.",
			},
			"token": {
				Type:        schema.TypeString,
				Computed:    true,
				Sensitive:   true,
				Description: "Generated runner token issued at runner registration.",
			},
			"runner_id": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "ID of the registered runner.",
			},
			"access_level": {
				Type:         schema.TypeString,
				ForceNew:     true,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"not_protected", "ref_protected"}, true),
				Description:  "Run against all refs, or protected only. Legal values are 'not_protected' or 'ref_protected'.",
			},
			"locked": {
				Type:        schema.TypeBool,
				ForceNew:    true,
				Optional:    true,
				Default:     true,
				Description: "Lock runner to project.",
			},
			"maximum_timeout": {
				Type:         schema.TypeInt,
				ForceNew:     true,
				Optional:     true,
				ValidateFunc: validation.IntAtLeast(10 * 60),
				Description:  "Maximum timeout for jobs.",
			},
			"tags": {
				Type:        schema.TypeSet,
				Optional:    true,
				ForceNew:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Set:         schema.HashString,
				Description: "List of tags for the runner.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Default:     "",
				Description: "Runner description; trivially visible from the UI.",
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
			"maintenance_note": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				Description:  "Free-form maintenance notes for the runner (255 characters max).",
				ValidateFunc: validation.StringLenBetween(0, 255),
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Runner 'name'.  Not highly visible in the UI; you probably want 'description'.",
			},

			"triggers": {
				Description: "A map of arbitrary strings that, when changed, will force the runner to be re-registered (and a new token issued).",
				Type:        schema.TypeMap,
				Optional:    true,
				ForceNew:    true,
			},
		},
	}

	for _, a := range s.Schema {
		if a.ForceNew {
			a.Description += fnDesc
		}
	}

	return s
}

func resourceGitlabRunnerCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	api := meta.(apiClient)

	tflog.Debug(ctx, "create gitlab runner token")

	// config.RegisterRunnerResponse omits the ID field :\
	type registrationResponse struct {
		config.RegisterRunnerResponse
		ID int `json:"id"`
	}

	query := config.RegisterRunnerRequest{
		Token: d.Get("registration_token").(string),
		RegisterRunnerParameters: config.RegisterRunnerParameters{
			Description:     d.Get("description").(string),
			RunUntagged:     d.Get("run_untagged").(bool),
			Active:          d.Get("active").(bool),
			Locked:          d.Get("locked").(bool),
			AccessLevel:     d.Get("access_level").(string),
			MaximumTimeout:  d.Get("maximum_timeout").(int),
			MaintenanceNote: d.Get("maintenance_note").(string),
		},
		Info: config.VersionInfo{
			Name:    d.Get("name").(string),
			Version: api.userAgent, // overwritten when the runner connects
			// Revision: "a revision",
		},
	}

	if v, ok := d.GetOk("tags"); ok {
		query.Tags = strings.Join(*stringSetToStringSlice(v.(*schema.Set)), ",")
	}

	url := api.baseURL + "/runners"

	j, err := json.Marshal(query)
	if err != nil {
		return diag.FromErr(err)
	}
	tflog.Trace(ctx, fmt.Sprintf("create gitlab runner query: %s", j))

	req := api.newAgent().Post(url).Send(query)

	// TODO other registration options...

	var runnerDetails registrationResponse
	resp, raw, errs := req.EndStruct(&runnerDetails)

	tflog.Trace(ctx, fmt.Sprintf("create gitlab runner token response: %s", raw))

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

	tflog.Debug(ctx, fmt.Sprintf("create gitlab runner token successful for runner #%d", runnerDetails.ID))
	return nil
}

func resourceGitlabRunnerRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	api := meta.(apiClient)

	tflog.Debug(ctx, fmt.Sprintf("validating runner token for runner #%s", d.Id()))

	runnerID, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	tflog.Trace(ctx, fmt.Sprintf("read gitlab runner %d", runnerID))

	url := api.baseURL + "/runners/verify"
	query := "token=" + d.Get("token").(string)
	req := api.newAgent().Post(url).Send(query)
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
	api := meta.(apiClient)

	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	tflog.Debug(ctx, fmt.Sprintf("Delete gitlab runner %d", id))

	url := api.baseURL + "/runners"
	query := "token=" + d.Get("token").(string)
	req := api.newAgent().Delete(url).Send(query)
	resp, _, errs := req.End()

	for _, err := range errs {
		// FIXME
		return diag.FromErr(err)
	}

	if resp.StatusCode == 204 {
		// all good!
		return nil
	}

	return diag.Errorf("bad response (%d): %s", resp.StatusCode, resp.Status)
}

func resourceGitlabRunnerUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	tflog.Debug(ctx, fmt.Sprintf("Update gitlab runner token %s", d.Id()))

	// GNDN -- the only thing we can update/change here without forcing a new
	// resource is the registration token -- and we're not taking any action
	// off of that.

	return nil
}
