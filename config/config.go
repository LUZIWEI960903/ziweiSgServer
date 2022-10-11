package config

import (
	"errors"
	"log"
	"os"

	"github.com/Unknwon/goconfig"
)

const configFile = "/conf/conf.ini"

var File *goconfig.ConfigFile

func init() {
	currentDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	configPath := currentDir + configFile

	l := len(os.Args)
	if l > 1 {
		dir := os.Args[1]
		if dir != "" {
			configPath = dir + configFile
		}
	}

	if !fileExist(configPath) {
		panic(errors.New("`conf.ini` is not exist!!"))
	}

	File, err = goconfig.LoadConfigFile(configPath)
	if err != nil {
		log.Fatalf("Loading `conf.ini` file error:", err)
	}
}

func fileExist(fileName string) bool {
	_, err := os.Stat(fileName)
	return err == nil || os.IsExist(err)
}
