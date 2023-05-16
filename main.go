package main

import (
	"database/sql"
	_ "fmt"
	"github.com/kelseyhightower/envconfig"
	"github.com/spf13/cobra"
	"log"
	"pr_ramadhan/database"
	"pr_ramadhan/handlers"
	"pr_ramadhan/models"
)

var rootCmd = &cobra.Command{Use: "app"}
var cfg models.Config

// createCmd represents the `create` command

func main() {
	// Initialize the Cobra CLI

	// Connect to the database

	// pr
	err1 := envconfig.Process("", &cfg)
	if err1 != nil {
		log.Fatal("error", err1)
	}

	// INITAL DATABASE
	Db, err := database.ConnnectDb(&cfg)
	if err != nil {
		log.Fatal(err)
	}

	// createCmd represents the `create` command
	var createCmd = &cobra.Command{
		Use:   "create",
		Short: "Create a new wiki entry",
		Run:   handlers.CreateWikiHandler(database.NewWikiRepository(Db)),
	}

	rootCmd.AddCommand(createCmd)

	// Execute the root command
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err, "error di executenya ")
	}

	// Close the database connection
	sqlDb, err := Db.DB()
	if err != nil {
		log.Fatal(err)
	}
	defer func(sqlDb *sql.DB) {
		err := sqlDb.Close()
		if err != nil {

		}
	}(sqlDb)
}

// Initialize the Cobra CLI
//func init() {
//	rootCmd.AddCommand(CreateCmd)
//}
