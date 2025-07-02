package main

import (
	"wallet/cmd"
	"wallet/helpers"
)

func main() {
	helpers.SetUpConfig()

	// load logger
	helpers.SetUpLogger()

	// connect to MySQL database
	helpers.SetUpMySQL()

	// start gRPC server
	go cmd.ServeGRPC()
	// start HTTP server
	cmd.ServeHTTP()
}
