package main

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
	"pr_ramadhan/models"
	"time"
)

// define the database connection
var db *gorm.DB

//func init() {
//	// connect to the database
//	dsn := "root:Pastibisa@tcp(localhost:3306)/pr_ramadhan"
//	// charset=utf8mb4&parseTime=True&loc=Local
//	var err error
//	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
//	if err != nil {
//		panic("failed to connect to database")
//	}
//
//	// run auto migration to create the table
//	err = db.AutoMigrate(&models.Wiki{})
//	if err != nil {
//		panic("failed to migrate database")
//	}
//}

// createCmd represents the `create` command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new wiki entry",
	Run: func(cmd *cobra.Command, args []string) {
		// prompt the user for the topic
		prompt := promptui.Prompt{
			Label: "Topic",
		}
		topic, err := prompt.Run()
		if err != nil {
			fmt.Println("Failed to read input")
			return
		}

		// create a new Wiki instance with the topic and timestamps
		now := time.Now()
		wiki := models.Wiki{
			Topic:     topic,
			CreatedAt: now,
			UpdatedAt: now,
		}

		// save the new wiki entry to the database
		err = db.Create(&wiki).Error
		if err != nil {
			fmt.Println("Failed to save data to database")
			return
		}

		// print a success message
		fmt.Println("Wiki terbaru di buat dengan ID", wiki.ID)
		// New wiki entry created with ID:
	},
}

func main() {
	// create a new root command
	var rootCmd = &cobra.Command{Use: "app"}
	// add the `create` command as a subcommand
	rootCmd.AddCommand(createCmd)

	// execute the root command
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
