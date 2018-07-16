package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	Name string `yaml:"name"`
	Collections []CollectionConfig `yaml:"collections"`
}

type CollectionConfig struct {
	Collection string `yaml:"collection"`
	Query string `yaml:"query"`
	Maps []MapConfig `yaml:"maps"`
}

type MapConfig struct {
	Name string `yaml:"name"`
	Key string `yaml:"key"`
	Value string `yaml:"val"`
}

func Load(path string) (Config, error) {
	raw, err := ioutil.ReadFile(path)
	if err != nil {
		return Config{}, err
	}

	var conf Config
	if err := yaml.Unmarshal(raw, &conf); err != nil {
		return Config{}, err
	}

	return conf, nil
}