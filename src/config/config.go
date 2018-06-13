package config

import (
	"log"
	"github.com/BurntSushi/toml"
)


type Config struct {
	host string
	db_index int
	db_orgReference int
	db_instReference int
	db_saleReference int
}

func (c *Config) Load() {
	if _, err := toml.DecodeFile("config.toml", &c); err != nil {
		log.Fatal(err)
	}
}