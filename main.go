package main

import (
	"knife-panel/server/core"
	"knife-panel/server/initialize"
)

func main() {
	initialize.InitConfig()
	initialize.InitLog()
	initialize.InitDatabase()
	core.RunServer()
}
