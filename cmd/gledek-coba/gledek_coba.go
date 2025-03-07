package main

import (
	"gledekcoba/internal/bootstrap"

	"github.com/sev-2/raiden"
)

func main() {
	// load configuration
	config, err := raiden.LoadConfig(nil)
	if err != nil {
		raiden.Panic(err.Error())
	}

	// Setup server
	server := raiden.NewServer(config)

	// register route
	bootstrap.RegisterRoute(server)

	// run server
	server.Run()
}
