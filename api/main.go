/**
 * Author: Goddy <goddy@mykg.ai> 2020-12-10
 */
package main

import (
	"fmt"
	"mykg.ai/rua/config"
	"mykg.ai/rua/web"
	"os"
)

func init() {
	config.Setup()
}

func main()  {
	env := os.Getenv("env")
	var path string
	if env == "" {
		path = "hi"
	} else {
		path = "yo"
	}
	fmt.Println("====")
	fmt.Print(path)


	web.Setup()
}
