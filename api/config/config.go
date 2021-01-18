/**
 * Author: Goddy <goddy@mykg.ai> 2020-12-10
 */
package config

import (
	"flag"
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
)

var (
	Env    string
	Config config
)

type config struct {
	Server struct {
		Port string
	} `yaml:",flow"`
	DB struct {
		Host     string
		Port     string
		Database string
		User     string
		Password string
	} `yaml:"db,flow"`
}

func Setup() {
	env := flag.String("env", "local", "运行环境")
	flag.Parse()
	Env = *env
	ymlBytes, err := ioutil.ReadFile(fmt.Sprintf("./config/%s.yml", *env))
	if err != nil {
		panic(err)
	}
	if err := yaml.Unmarshal([]byte(os.ExpandEnv(string(ymlBytes))), &Config); err != nil {
		panic(err)
	}

	connectDB()
	migrateDomains()
}
