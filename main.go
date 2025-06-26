package main

import (
	"ewallet-framework/cmd"
	"ewallet-framework/helpers"
)

func main() {
	// Load configuration
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
