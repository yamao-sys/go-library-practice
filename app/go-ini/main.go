package main

import (
	"fmt"

	"gopkg.in/go-ini/ini.v1"
)

type ConfigList struct {
	Port      int
	DbName    string
	SQLDriver string
}

var Config ConfigList

func init() {
	cfg, _ := ini.Load("config.ini")

	Config = ConfigList{
		Port:   cfg.Section("web").Key("port").MustInt(8080),
		DbName: cfg.Section("db").Key("name").MustString("go_library_practice"),
	}
}

func main() {
	fmt.Printf("Port = %v\n", Config.Port)
	fmt.Printf("DbName = %v\n", Config.DbName)
}
