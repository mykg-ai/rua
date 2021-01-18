/**
 * Author: Goddy <goddy@mykg.ai> 2020-12-10
 */
package main

import (
	"mykg.ai/rua/config"
	"mykg.ai/rua/web"
)

func init() {
	config.Setup()
}

func main()  {
	web.Setup()
}
