package main

import (
	"clipboard/client"
	"clipboard/server"
	"flag"
)

func main() {
	mode := flag.String("m", "server", "mode to activate; server or client")
	flag.Parse()

	switch *mode {
	case "server":
		server.RunServer()
	case "client":
		client.RunClient()
	}
}
