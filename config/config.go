package config

import (
	"github.com/kelseyhightower/envconfig"
	"log"
	)

type MySqlConfiguration struct {
	Host     string `envconfig:"MYSQL_HOST"`
	User     string `envconfig:"MYSQL_USER"`
	Password string `envconfig:"MYSQL_PASSWORD"`
	Database string `envconfig:"MYSQL_DATABASE"`
}

type Configuration struct {
	MySQL MySqlConfiguration `json:"mysql"`
}

var Config Configuration

func Configure() {
	mysql := &MySqlConfiguration{}
	mysqlParseError := envconfig.Process("mysql", mysql)

	if mysqlParseError != nil {
		log.Fatal(mysqlParseError.Error())
	}

	Config = Configuration{
		MySQL: *mysql,
	}
}
