package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type RequestParameters struct {
	// TODO support multiple values per key
	QueryParameters map[string]string `yaml:"query_params"`
	// TODO support arbitrary json
	JSONParameters map[string]string `yaml:"json_params"`
	Headers        map[string]string `yaml:"headers"`
}

type GaghConfig struct {
	Port             string            `yaml:"port"`
	ConfigParameters map[string]string `yaml:"config_parameters"`

	WebhookConfigs []struct {
		RequestParameters
		Name string `yaml:"name"`
		Url  string `yaml:"url"`
	} `yaml:"webhook_configs"`

	// TODO don't handle each type of request independently
	GitHubPullRequestHandlers []struct {
		// TODO allow response based on more generic criteria than just the action
		Action string `yaml:"action"`

		Functions []struct {
			Name      string   `yaml:"name"`
			Arguments []string `yaml:"arguments"`
		}

		Webhooks []struct {
			RequestParameters
			Name string `yaml:"name"`
		} `yaml:"webhooks"`
	}
}

func LoadConfig() (*GaghConfig, error) {
	config_bytes, err := ioutil.ReadFile("/etc/gagh/config.yml")
	if err != nil {
		return nil, err
	}

	conf := &GaghConfig{}
	if err := yaml.Unmarshal([]byte(config_bytes), &conf); err != nil {
		return nil, err
	}
	return conf, nil
}
