package main

import (
	"fmt"
	"github.com/open-dojo-kubernetes/scoreboard/service"
	"github.com/open-dojo-kubernetes/scoreboard/dbclient"
)

var appName = "scoreboard"

func main() {
	fmt.Printf("Starting %v\n", appName)
	initializeBoltClient()
	service.StartWebServer("6767")
}

// Creates instance and calls the OpenBoltDb and Seed funcs
func initializeBoltClient() {
	service.DBClient = &dbclient.BoltClient{}
	service.DBClient.OpenBoltDb()
	service.DBClient.StartAGame()
}