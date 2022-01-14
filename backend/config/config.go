package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/tkanos/gonfig"
)

type Config struct {
	DB_PORT     int
	DB_HOST     string
	DB_Username string
	DB_Pass     string
	DB_Name     string
	DB_Table    []string
}

var err error

func GetConfig() Config {
	conf := Config{}
	gonfig.GetConf("config/config.json", &conf)
	if len(os.Getenv("DB_PORT")) != 0 {
		conf.DB_PORT, err = strconv.Atoi(os.Getenv("DB_PORT"))
		if err != nil {
			fmt.Println("Can't get DB_PORT from environment, reverting to port", 3306)
			conf.DB_PORT = 3306
		}
	}
	conf.DB_Name = "DB"
	if len(os.Getenv("DB_HOST")) != 0 {
		conf.DB_HOST = os.Getenv("DB_HOST")
	}
	if len(os.Getenv("DB_Username")) != 0 {
		conf.DB_Username = os.Getenv("DB_Username")
	}
	if len(os.Getenv("DB_Pass")) != 0 {
		conf.DB_Pass = os.Getenv("DB_Pass")
	}
	if len(os.Getenv("DB_Name")) != 0 {
		conf.DB_Name = os.Getenv("DB_Name")
	}

	return conf
}
