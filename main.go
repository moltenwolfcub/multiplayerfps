package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/moltenwolfcub/multiplayerfps/client"
	"github.com/moltenwolfcub/multiplayerfps/common"
	"github.com/moltenwolfcub/multiplayerfps/server"
)

var sideFlag string
var portFlag string

func main() {
	common.RegisterPackets()

	flag.StringVar(&sideFlag, "side", "", "'server' or 'client'")
	flag.StringVar(&portFlag, "port", "", "desired port to connect to or host from")
	flag.Parse()

	if sideFlag == "server" {
		if sideFlag == "" {
			sideFlag = ":0"
		}

		fmt.Println("server")

		server := server.NewServer(portFlag)
		cleanup := common.SetupServerLogger()
		defer cleanup()

		server.Start()

	} else if sideFlag == "client" {
		if sideFlag == "" {
			log.Fatal("Please specify a port")
		}

		fmt.Println("client")
		client := client.NewClient(portFlag)

		log.Fatal(client.Start())
	} else {
		log.Fatal("Unknown side to launch")
	}
}
