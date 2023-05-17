package handlers

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"pr_ramadhan/models"
	"pr_ramadhan/repoWiki"
	"strconv"
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
		wiki := models.Wikis{
			Topic:     topic,
			CreatedAt: now.Format(time.DateTime),
			UpdatedAt: now.Format(time.DateTime),
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

func UpdateWikiHandler(repo repoWiki.WikiRepository) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		// meminta pengguna untuk memasukkan ID wiki yang akan diperbarui
		prompt := promptui.Prompt{
			Label: "Wiki ID",
		}
		idStr, err := prompt.Run()
		if err != nil {
			fmt.Println("Failed to read input")
			return
		}

		// mengkonversi ID menjadi tipe uint
		id, err := strconv.ParseUint(idStr, 10, 32)
		if err != nil {
			fmt.Println("Invalid ID")
			return
		}

		// meminta pengguna untuk memasukkan topik baru
		topicPrompt := promptui.Prompt{
			Label: "New Topic",
		}
		newTopic, err := topicPrompt.Run()
		if err != nil {
			fmt.Println("Failed to read input")
			return
		}

		// membangun objek Wiki yang akan diperbarui
		wiki := &models.Wikis{
			ID:    uint(id),
			Topic: newTopic,
		}

		// memanggil fungsi UpdateWiki pada repository
		err = repo.UpdateWiki(wiki)
		if err != nil {
			fmt.Println("Failed to update wiki")
			return
		}

		// print a success message
		fmt.Println("Wiki with ID", id, "has been updated")
	}
}

func GetWikiHandler(repo repoWiki.WikiRepository) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		// meminta pengguna untuk memasukkan ID wiki
		prompt := promptui.Prompt{
			Label: "Wiki ID",
		}
		idStr, err := prompt.Run()
		if err != nil {
			fmt.Println("Failed to read input")
			return
		}

		// mengkonversi ID menjadi tipe uint
		id, err := strconv.ParseUint(idStr, 10, 32)
		if err != nil {
			fmt.Println("Invalid ID")
			return
		}

		// memanggil fungsi GetWiki pada repository
		wiki, err := repo.GetWiki(uint(id))
		if err != nil {
			fmt.Println("Failed to get wiki from database / Not Found")
			return
		}

		// menampilkan hasil
		fmt.Println("Wiki ID:", wiki.ID)
		fmt.Println("Topic:", wiki.Topic)
		fmt.Println("Description:", wiki.Description)
		fmt.Println("Created At:", wiki.CreatedAt)
		fmt.Println("Updated At:", wiki.UpdatedAt)
	}
}

func DeleteWikiHandler(repo repoWiki.WikiRepository) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		// meminta pengguna untuk memasukkan ID wiki yang akan dihapus
		prompt := promptui.Prompt{
			Label: "Wiki ID",
		}
		idStr, err := prompt.Run()
		if err != nil {
			fmt.Println("Failed to read input")
			return
		}

		// mengkonversi ID menjadi tipe uint
		id, err := strconv.ParseUint(idStr, 10, 32)
		if err != nil {
			fmt.Println("Invalid ID")
			return
		}

		// memanggil fungsi DeleteWiki pada repository
		err = repo.DeleteWiki(uint(id))
		if err != nil {
			fmt.Println("Failed to delete wiki from database")
			return
		}

		// print a success message
		fmt.Println("Wiki with ID", id, "has been deleted")
	}
}
