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

	if !fileExist(configPath) {
		panic(errors.New("`conf.ini` is not exist!!"))
	}

	l := len(os.Args)
	if l > 1 {
		dir := os.Args[1]
		configPath = dir + configFile
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
