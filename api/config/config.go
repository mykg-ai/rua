/**
 * Author: Goddy <goddy@mykg.ai> 2020-12-10
 */
package config

import "C"
import (
	"flag"
	"fmt"
	"gopkg.in/ini.v1"
	"os"
)

var (
	Env string
	Server server
	dbConfig db
)

type server struct {
	HttpPort string
}

type db struct {
	Host string
	Port string
	Database string
	User string
	Password string
}

/** === config === */

func Setup() {
	env := flag.String("env", "local", "运行环境")
	flag.Parse()
	Env = *env
	cfg, err := ini.Load(fmt.Sprintf("./config/%s.ini", *env))
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
	toBean(cfg, "server", &Server)
	toBean(cfg, "db", &dbConfig)

	ConnectDB()
	MigrateDomains()
}

func toBean(cfg *ini.File, section string, bean interface{}) {
	err := cfg.Section(section).MapTo(bean)
	if err != nil {
		fmt.Printf("ini map %s error: %v", section, err)
	}
}