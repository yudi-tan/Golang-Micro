package main

import (
	"fmt"
	"github.com/yuditan/goblog/accountservice/dbclient"
	"github.com/yuditan/goblog/accountservice/service"
)

var appName = "accountservice"

func main() {
	fmt.Printf("Starting %v\n", appName)
	initializeBoltClient()
	service.StartWebServer("6767")
}

func initializeBoltClient() {
	service.DBclient = &dbclient.BoltClient{}
	service.DBclient.OpenBoltDB()
	service.DBclient.Seed()
}