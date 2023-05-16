package handlers

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"pr_ramadhan/models"
	"pr_ramadhan/repoWiki"
	"time"
)

func CreateWikiHandler(repo repoWiki.WikiRepository) func(cmd *cobra.Command, args []string) {

	return func(cmd *cobra.Command, args []string) {
		// meminta pengguna untuk menentukan topik
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
		err = repo.AddWiki(&wiki)
		if err != nil {
			fmt.Println("Failed to save data to database")
			return
		}

		// print a success message
		fmt.Println("Wiki terbaru di buat dengan ID", wiki.ID)
	}
}
