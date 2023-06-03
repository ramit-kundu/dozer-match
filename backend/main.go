package main

import (
	"log"
	"os"

	"github.com/kundu-ramit/dozer_match/cmd"
	"github.com/kundu-ramit/dozer_match/infra/database"
	"github.com/kundu-ramit/dozer_match/infra/redis"

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
	case "server":
		startServer()
	case "migration":
		applyMigration()
	default:
		log.Fatal("Invalid command:", args[0])
	}
}

func applyMigration() {

	db := database.Initialize()
	cmd.ApplyMigration(db)

}

func startServer() {

	database.Initialize()
	redis.Initialize()
	r := routes.SetupRouter()
	r.Run(":8002")
}
