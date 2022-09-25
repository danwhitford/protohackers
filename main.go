package main

import (
	"log"
	"net"
	"os"

	"github.com/danwhitford/protohackers/budgetchat"
	"github.com/danwhitford/protohackers/meanstoanend"
	"github.com/danwhitford/protohackers/primetime"
	"github.com/danwhitford/protohackers/simpleserver"
	"github.com/danwhitford/protohackers/smoketest"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalln("Need to give command")
	}

	funcs := map[string]func(c net.Conn){
		"smoketest":    smoketest.HandleFunc,
		"primetime":    primetime.HandleFunc,
		"meanstoanend": meanstoanend.HandleFunc,
	}
	cmd := os.Args[1]

	switch cmd {
	case "smoketest", "primetime", "meanstoanend":
		var handlerFunc func(c net.Conn)
		handlerFunc = funcs[cmd]
		simpleserver.SimpleServer(handlerFunc)
	case "budgetchat":
		budgetchat.ChatServer()
	default:
		log.Fatalf("Command not found: '%s'", cmd)
	}
}
