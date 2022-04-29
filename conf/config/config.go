package config

import (
	"log"
	"os"

	"gopkg.in/ini.v1"
)

type ConfigList struct {
	Mysql  MysqlConfig
	Redis  RedisConfig
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

type RedisConfig struct {
	RedisHost string
	RedisPort string
}

type ServerConfig struct {
	ServerPort int
	MasterKey  string
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

	redis := RedisConfig{
		RedisHost: cfg.Section("redis").Key("redis_host").String(),
		RedisPort: cfg.Section("redis").Key("redis_port").String(),
	}

	server := ServerConfig{
		ServerPort: cfg.Section("server").Key("server_port").MustInt(),
		MasterKey:  cfg.Section("server").Key("master_key").String(),
	}

	Config = ConfigList{
		Mysql:  mysql,
		Server: server,
		Redis:  redis,
	}
}
