package main

import (
	"context"
	"log"
	"os"

	"github.com/kundu-ramit/dozer_match/cmd"
	"github.com/kundu-ramit/dozer_match/database"
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
	case "migration":
		applyMigration()
	default:
		log.Fatal("Invalid command:", args[0])
	}
}

func startCron() {
}

func applyMigration() {

	db, err := database.InitializeDatabase()
	if err != nil {
		panic(err)
	}
	cmd.ApplyMigration(context.TODO(), db)

}

func startServer() {

	r := routes.SetupRouter()
	r.Run(":8002")
}
