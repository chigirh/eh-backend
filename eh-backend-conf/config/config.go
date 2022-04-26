package config

import (
	"log"
	"os"

	"gopkg.in/ini.v1"
)

type ConfigList struct {
	Mysql  MysqlConfig
	Server ServerConfig
}

type MysqlConfig struct {
	DbDriverName   string
	DbName         string
	DbUserName     string
	DbUserPassword string
	DbHost         string
	DbPort         string
}

type ServerConfig struct {
	ServerPort int
	ApiKey     string
}

var Config ConfigList

func init() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Printf("Failed to read file: %v", err)
		os.Exit(1)
	}

	mysql := MysqlConfig{
		DbDriverName:   cfg.Section("db").Key("db_driver_name").String(),
		DbName:         cfg.Section("db").Key("db_name").String(),
		DbUserName:     cfg.Section("db").Key("db_user_name").String(),
		DbUserPassword: cfg.Section("db").Key("db_user_password").String(),
		DbHost:         cfg.Section("db").Key("db_host").String(),
		DbPort:         cfg.Section("db").Key("db_port").String(),
	}

	server := ServerConfig{
		ServerPort: cfg.Section("server").Key("server_port").MustInt(),
		ApiKey:     cfg.Section("server").Key("api_key").String(),
	}

	Config = ConfigList{
		Mysql:  mysql,
		Server: server,
	}
}
