// This file is largely snagged from the gitlab-runner repo, as the
// structs/etc we're interested in are unexported :\

package configtemplate

import (
	"github.com/imdario/mergo"
	"github.com/pkg/errors"
	"gitlab.com/gitlab-org/gitlab-runner/common"
)

type ConfigTemplate struct {
	*common.Config

	configFile string `long:"config" env:"TEMPLATE_CONFIG_FILE" description:"Path to the configuration template file"`
}

func (c *ConfigTemplate) Enabled() bool {
	return c.configFile != ""
}

func (c *ConfigTemplate) MergeTo(config *common.RunnerConfig) error {
	err := c.LoadConfigTemplate()
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

func NewConfigTemplateFromFile(cfgFile string) (*ConfigTemplate, error) {
	ct := ConfigTemplate{
		configFile: cfgFile,
	}
	if err := ct.LoadConfigTemplate(); err != nil {
		return nil, err
	}
	return &ct, nil
}

// func NewConfigTemplateFromString(cfg string) (*ConfigTemplate, error) {
// 	// https://godoc.org/github.com/hashicorp/terraform-plugin-sdk/helper/pathorcontents
// 	templateStr := pathorcontents.Read(templateData)
// 	tmpfile, err := ioutil.TempFile("", "runnercfgtemplate*")
// 	if err != nil {
// 		return err
// 	}
// 	defer os.Remove(tmpfile.Name())
// 	if _, err = tmpfile.WriteContent(templateStr); err != nil {
// 		return err
// 	}
// 	ct := ConfigTemplate{
// 		configFile: cfgFile,
// 	}
// 	if err := ct.LoadConfigTemplate(); err != nil {
// 		return nil, err
// 	}
// 	return ct, nil
// }

func (c *ConfigTemplate) LoadConfigTemplate() error {
	config := common.NewConfig()

	err := config.LoadConfig(c.configFile)
	if err != nil {
		return err
	}

	c.Config = config

	return nil
}
