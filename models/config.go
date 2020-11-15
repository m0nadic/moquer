package models

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io"
	"io/ioutil"
)

type Config struct {
	Version  string
	Services map[string]*Service
}

func (cfg Config) PrintYaml(w io.Writer) error {
	data, err := yaml.Marshal(cfg)
	if err != nil {
		return err
	}
	_, _ = fmt.Fprintln(w, string(data))
	return nil
}

func NewConfigFromFile(fileName string) (*Config, error) {
	data, err := ioutil.ReadFile(fileName)

	if err != nil {
		return nil, err
	}
	var cfg Config
	err = yaml.Unmarshal(data, &cfg)

	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
