package main

import (
	"log"
	"os"

	"github.com/kundu-ramit/dozer_match/routes"
)

func main() {
	// Get the command-line arguments
	args := os.Args[1:]

	// Check the number of arguments
	if len(args) == 0 {
		log.Fatal("No command specified.")
	}

	// Handle the command
	switch args[0] {
	case "cron":
		startCron()
	case "server":
		startServer()
	default:
		log.Fatal("Invalid command:", args[0])
	}
}

func startCron() {
}

func startServer() {

	r := routes.SetupRouter()
	r.Run(":8002")
}
