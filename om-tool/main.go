package main

import (
	cache "om-tool/src/cache"
	config "om-tool/src/config"
	"os"
	"flag"
	"fmt"
)

var (
	defaultMongoInfo = "mongodb://eagle_app_user:eagleappuser@ftc-lbeagmdb306:27017,ftc-lbeagmdb307:27017,ftc-lbeagmdb308:27017/ODS"
	defaultPath = "./config.yml"
)

var (
	configPath string
	mongoInfo string
)

func init() {
	if envVal := os.Getenv("MONGO_DB"); envVal != "" {
		defaultMongoInfo = envVal
	}
	flag.StringVar(&configPath, "c", defaultPath, "config path")
	flag.StringVar(&mongoInfo, "m", defaultMongoInfo, "mongo url")
}

func main() {
	flag.Parse()
	conf, _ := config.Load(configPath)
	if err := cache.SetupCache(conf, mongoInfo); err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
}
