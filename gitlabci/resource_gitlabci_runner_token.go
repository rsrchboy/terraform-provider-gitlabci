package gitlabci

// https://docs.gitlab.com/ce/api/runners.html#register-a-new-runner

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/parnurzeal/gorequest"
)

func resourceGitlabRunner() *schema.Resource {
	return &schema.Resource{
		Create: resourceGitlabRunnerCreate,
		Read:   resourceGitlabRunnerRead,
		Delete: resourceGitlabRunnerDelete,

		Schema: map[string]*schema.Schema{
			"registration_token": {
				Type:      schema.TypeString,
				ForceNew:  true,
				Required:  true,
				Sensitive: true,
				StateFunc: hashSum,
			},
			"token": {
				Type:      schema.TypeString,
				Computed:  true,
				Sensitive: true,
			},
			"runner_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"access_level": {
				Type:         schema.TypeString,
				ForceNew:     true,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"not_protected", "ref_protected"}, true),
			},
			"locked": {
				Type:     schema.TypeBool,
				ForceNew: true,
				Optional: true,
				Default:  true,
			},
			"maximum_timeout": {
				Type:         schema.TypeInt,
				ForceNew:     true,
				Optional:     true,
				ValidateFunc: validation.IntAtLeast(10 * 60),
			},
			"tags": {
				Type:     schema.TypeSet,
				Optional: true,
				ForceNew: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Default:  "",
			},
			"active": {
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: true,
				Default:  true,
			},
			"run_untagged": {
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: true,
				Default:  true,
			},
		},
	}
}

func resourceGitlabRunnerCreate(d *schema.ResourceData, meta interface{}) error {
	baseURL := meta.(string)

	log.Printf("[DEBUG] create gitlab runner token")

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
	log.Printf("[DEBUG] create gitlab runner query: %s", j)

	req := gorequest.
		New().
		Post(url).
		Send(query)

	// TODO other registration options...

	var runnerDetails registrationResponse
	resp, _, errs := req.EndStruct(&runnerDetails)

	for _, err := range errs {
		// FIXME
		return err
	}

	if resp.StatusCode != 201 {
		return errors.New(resp.Status)
	}

	d.SetId(fmt.Sprintf("%d", runnerDetails.ID))
	d.Set("token", runnerDetails.Token)
	d.Set("runner_id", runnerDetails.ID)

	return nil
}

func resourceGitlabRunnerRead(d *schema.ResourceData, meta interface{}) error {
	baseURL := meta.(string)

	runnerID, err := strconv.Atoi(d.Id())
	if err != nil {
		return err
	}

	// FIXME probably ought to VerifyRegisteredRunner() here first

	log.Printf("[DEBUG] read gitlab runner %d", runnerID)

	url := baseURL + "/runners/verify"
	req := gorequest.
		New().
		Post(url).
		Send("token=" + d.Get("token").(string))
	resp, _, errs := req.End()

	for _, err := range errs {
		// FIXME
		return err
	}

	if resp.StatusCode == 200 {
		// all good!
		return nil
	}

	return errors.New(resp.Status)
}

func resourceGitlabRunnerDelete(d *schema.ResourceData, meta interface{}) error {
	baseURL := meta.(string)

	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Delete gitlab runner %d", id)

	url := baseURL + "/runners"
	req := gorequest.
		New().
		Delete(url).
		Send("token=" + d.Get("token").(string))
	resp, _, errs := req.End()

	for _, err := range errs {
		// FIXME
		return err
	}

	if resp.StatusCode == 204 {
		// all good!
		return nil
	}

	return errors.New(resp.Status)
}
