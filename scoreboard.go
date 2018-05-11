package main

import (
	"fmt"
	"github.com/lPegz/kubernetes-ping-pong/scoreboard/service"
	"github.com/lPegz/kubernetes-ping-pong/scoreboard/dbclient"
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