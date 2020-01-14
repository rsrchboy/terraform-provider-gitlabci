// This file is largely snagged from the gitlab-runner repo, as the
// structs/etc we're interested in are unexported :\

package configtemplate

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"strings"

	"github.com/imdario/mergo"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"

	"gitlab.com/gitlab-org/gitlab-runner/common"
)

type ConfigTemplate struct {
	*common.Config

	ConfigFile string `long:"config" env:"TEMPLATE_CONFIG_FILE" description:"Path to the configuration template file"`
}

func (c *ConfigTemplate) Enabled() bool {
	return c.ConfigFile != ""
}

func (c *ConfigTemplate) MergeTo(config *common.RunnerConfig) error {
	err := c.loadConfigTemplate()
	if err != nil {
		return errors.Wrap(err, "couldn't load configuration template file")
	}

	if len(c.Runners) != 1 {
		return errors.New("configuration template must contain exactly one [[runners]] entry")
	}

	err = mergo.Merge(config, c.Runners[0])
	if err != nil {
		return errors.Wrap(err, "error while merging configuration with configuration template")
	}

	return nil
}

func (c *ConfigTemplate) loadConfigTemplate() error {
	config := common.NewConfig()

	err := config.LoadConfig(c.ConfigFile)
	if err != nil {
		return err
	}

	c.Config = config

	return nil
}

type RegisterCommand struct {
	context    *cli.Context
	network    common.Network
	reader     *bufio.Reader
	registered bool

	configOptions

	ConfigTemplate ConfigTemplate `namespace:"template"`

	TagList           string `long:"tag-list" env:"RUNNER_TAG_LIST" description:"Tag list"`
	NonInteractive    bool   `short:"n" long:"non-interactive" env:"REGISTER_NON_INTERACTIVE" description:"Run registration unattended"`
	LeaveRunner       bool   `long:"leave-runner" env:"REGISTER_LEAVE_RUNNER" description:"Don't remove runner if registration fails"`
	RegistrationToken string `short:"r" long:"registration-token" env:"REGISTRATION_TOKEN" description:"Runner's registration token"`
	RunUntagged       bool   `long:"run-untagged" env:"REGISTER_RUN_UNTAGGED" description:"Register to run untagged builds; defaults to 'true' when 'tag-list' is empty"`
	Locked            bool   `long:"locked" env:"REGISTER_LOCKED" description:"Lock Runner for current project, defaults to 'true'"`
	AccessLevel       string `long:"access-level" env:"REGISTER_ACCESS_LEVEL" description:"Set access_level of the runner to not_protected or ref_protected; defaults to not_protected"`
	MaximumTimeout    int    `long:"maximum-timeout" env:"REGISTER_MAXIMUM_TIMEOUT" description:"What is the maximum timeout (in seconds) that will be set for job when using this Runner"`
	Paused            bool   `long:"paused" env:"REGISTER_PAUSED" description:"Set Runner to be paused, defaults to 'false'"`

	common.RunnerConfig
}

type AccessLevel string

const (
	NotProtected AccessLevel = "not_protected"
	RefProtected AccessLevel = "ref_protected"
)

const (
	defaultDockerWindowCacheDir = "c:\\cache"
)

// func (s *RegisterCommand) addRunner(runner *common.RunnerConfig) {
// 	s.config.Runners = append(s.config.Runners, runner)
// }

func (s *RegisterCommand) mergeTemplate() {
	if !s.ConfigTemplate.Enabled() {
		return
	}

	logrus.Infof("Merging configuration from template file %q", s.ConfigTemplate.ConfigFile)

	err := s.ConfigTemplate.MergeTo(&s.RunnerConfig)
	if err != nil {
		logrus.WithError(err).Fatal("Could not handle configuration merging from template file")
	}
}
