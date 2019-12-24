package gitlabci

// https://docs.gitlab.com/ce/api/runners.html#register-a-new-runner

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	// "strings"

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
				// StateFunc: hashSum,
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
				Computed:     true,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"not_protected", "ref_protected"}, true),
			},
			"locked": {
				Type:     schema.TypeBool,
				Computed: true,
				Optional: true,
			},
			"is_shared": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"maximum_timeout": {
				Type:         schema.TypeInt,
				Computed:     true,
				Optional:     true,
				ValidateFunc: validation.IntAtLeast(10 * 60),
			},
			"tags": {
				Type:     schema.TypeSet,
				Optional: true,
				// FIXME
				ForceNew: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				// FIXME
				ForceNew: true,
				Default:  "",
			},
			"active": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"run_untagged": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceGitlabRunnerCreate(d *schema.ResourceData, meta interface{}) error {
	baseURL := meta.(string)

	log.Printf("[DEBUG] create gitlab runner token")

	type RegisterOptions struct {
		Token          string   `json:"token"`
		Description    string   `json:"description"`
		Active         bool     `json:"active"`
		Locked         bool     `json:"locked"`
		RunUntagged    bool     `json:"run_untagged"`
		TagList        []string `json:"tag_list"`
		AccessLevel    string   `json:"access_level"`
		MaximumTimeout int      `json:"maximum_timeout"`
		// info hash
	}

	type registrationResponse struct {
		ID    int    `json:"id"`
		Token string `json:"token"`
	}

	registrationToken := d.Get("registration_token").(string)
	log.Printf("[DEBUG] create gitlab runner token - QWERTY HERE !!!!!!!!!!!!!!!!!1")
	log.Printf("[DEBUG] create gitlab runner token / token: " + registrationToken)
	url := baseURL + "/runners"
	log.Printf("[DEBUG] create gitlab runner token / url: " + url)
	req := gorequest.
		New().
		Post(url).
		Send("token=" + registrationToken)

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

	// // https://godoc.org/github.com/xanzy/go-gitlab#RegisterNewRunnerOptions
	// options := RegisterOptions{
	// 	token:        &registrationToken,
	// 	description:  d.Get("description").(string),
	// 	run_untagged: d.Get("run_untagged").(bool),
	// 	active:       d.Get("active").(bool),
	// 	locked:       d.Get("locked").(bool),
	// }
	// if v, ok := d.GetOk("tags"); ok {
	// 	options.TagList = *(stringSetToStringSlice(v.(*schema.Set)))
	// }

	// if v, ok := d.GetOk("maximum_timeout"); ok {
	// 	options.MaximumTimeout = gitlab.Int(v.(int))
	// }

	// runnerDetails, _, err := client.Runners.RegisterNewRunner(&options)
	// if err != nil {
	// 	return err
	// }

	d.SetId(fmt.Sprintf("%d", runnerDetails.ID))
	d.Set("token", runnerDetails.Token)

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

// func resourceGitlabRunnerUpdate(d *schema.ResourceData, meta interface{}) error {
// 	client := meta.(*gitlab.Client)
// 	id, err := strconv.Atoi(d.Id())
// 	if err != nil {
// 		return err
// 	}

// 	// https://godoc.org/github.com/xanzy/go-gitlab#UpdateRunnerDetailsOptions
// 	options := gitlab.UpdateRunnerDetailsOptions{
// 		Description: gitlab.String(d.Get("description").(string)),
// 		RunUntagged: gitlab.Bool(d.Get("run_untagged").(bool)),
// 		Active:      gitlab.Bool(d.Get("active").(bool)),
// 		Locked:      gitlab.Bool(d.Get("locked").(bool)),
// 		AccessLevel: gitlab.String(d.Get("access_level").(string)),
// 		// MaximumTimeout: gitlab.Int(d.Get("maximum_timeout").(int)),
// 		// X: gitlab.String(d.Get("X").(string)),
// 	}

// 	if v, ok := d.GetOk("tags"); ok {
// 		options.TagList = *(stringSetToStringSlice(v.(*schema.Set)))
// 	}

// 	if v, ok := d.GetOk("maximum_timeout"); ok {
// 		options.MaximumTimeout = gitlab.Int(v.(int))
// 	}

// 	log.Printf("[DEBUG] update gitlab runner %d", id)

// 	_, _, err = client.Runners.UpdateRunnerDetails(id, &options)
// 	if err != nil {
// 		return err
// 	}

// 	return resourceGitlabRunnerRead(d, meta)
// }

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
