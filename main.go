package main

import (
	"database/sql"
	_ "fmt"
	"github.com/kelseyhightower/envconfig"
	"github.com/spf13/cobra"
	"log"
	database2 "pr_ramadhan/cmd/database"
	"pr_ramadhan/cmd/handlers"
	"pr_ramadhan/cmd/models"
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
	Db, err := database2.ConnnectDb(&cfg)
	if err != nil {
		log.Fatal(err)
	}

	// createCmd represents the `create` command
	var createCmd = &cobra.Command{
		Use:   "create",
		Short: "Create a new wiki entry",
		Run:   handlers.CreateWikiHandler(database2.NewWikiRepository(Db)),
	}

	var updateCmd = &cobra.Command{
		Use:   "update",
		Short: "Update a wiki By Id",
		Run:   handlers.UpdateWikiHandler(database2.NewWikiRepository(Db)),
	}

	var getCmd = &cobra.Command{
		Use:   "get",
		Short: "Get a wiki By Id",
		Run:   handlers.GetWikiHandler(database2.NewWikiRepository(Db)),
	}

	var deleteCmd = &cobra.Command{
		Use:   "delete",
		Short: "Delete a Wiki By Id",
		Run:   handlers.DeleteWikiHandler(database2.NewWikiRepository(Db)),
	}

	//handlers.StartWorker(Db)

	//var workerCmd = &cobra.Command{
	//	Use:   "worker",
	//	Short: "Run the worker for scraping",
	//	Run:   handlers.StartWorker(Db),
	//}

	rootCmd.AddCommand(createCmd)
	rootCmd.AddCommand(updateCmd)
	rootCmd.AddCommand(getCmd)
	rootCmd.AddCommand(deleteCmd)
	//rootCmd.AddCommand(workerCmd)

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
