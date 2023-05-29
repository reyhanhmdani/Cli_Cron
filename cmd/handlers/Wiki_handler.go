package handlers

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"net/http"
	"pr_ramadhan/cmd/models"
	"pr_ramadhan/repoWiki"
	"strconv"
	"sync"
	"time"
)

// struct ini untuk melakukan scarping paragraf pertama dari Wiki
type Scrapper struct {
	Client *http.Client
}

// //////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func CreateWikiHandler(repo repoWiki.WikiRepository) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		// Meminta pengguna untuk menentukan topik
		prompt := promptui.Prompt{
			Label: "Topic",
		}
		topic, err := prompt.Run()
		if err != nil {
			fmt.Println("Failed to read input")
			return
		}

		// Membuat instance Wiki baru dengan topik dan timestamp
		now := time.Now()
		wiki := models.Wikis{
			Topic:       topic,
			CreatedAt:   now.Format("2006-01-02 15:04:05"),
			UpdatedAt:   now.Format("2006-01-02 15:04:05"),
			Description: "",
		}

		// Menyimpan entri wiki baru ke database
		err = repo.AddWiki(&wiki)
		if err != nil {
			fmt.Println("Failed to save data to database")
			return
		}

		// Menampilkan pesan sukses
		fmt.Println("Wiki terbaru dibuat dengan ID", wiki.ID)
	}
}

func UpdateWikiHandler(repo repoWiki.WikiRepository) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		// Meminta pengguna untuk memasukkan ID wiki yang akan diupdate
		prompt := promptui.Prompt{
			Label: "Wiki ID",
		}
		idStr, err := prompt.Run()
		if err != nil {
			fmt.Println("Failed to read input")
			return
		}

		id, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Println("Invalid ID")
			return
		}

		// Meminta pengguna untuk memasukkan topik yang baru
		prompt = promptui.Prompt{
			Label: "New Topic",
		}
		newTopic, err := prompt.Run()
		if err != nil {
			fmt.Println("Failed to read input")
			return
		}

		// Mengambil data wiki berdasarkan ID
		wiki, err := repo.GetWiki(id)
		if err != nil {
			fmt.Println("Failed to get data from database")
			return
		}

		// Mengupdate topik wiki
		wiki.Topic = newTopic

		// Mengatur zona waktu yang sesuai
		loc, err := time.LoadLocation("Asia/Jakarta")
		if err != nil {
			fmt.Println("Failed to load location")
			return
		}
		currentTime := time.Now().In(loc)
		wiki.UpdatedAt = currentTime.Format("2006-01-02 15:04:05")

		// Menyimpan perubahan ke database
		err = repo.UpdateWiki(wiki)
		if err != nil {
			fmt.Println("Failed to update data in database")
			return
		}

		// Menampilkan pesan sukses
		fmt.Println("Topic wiki dengan ID", id, "berhasil diupdate")
	}
}

func GetWikiHandler(repo repoWiki.WikiRepository) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		// meminta pengguna untuk memasukkan ID wiki
		prompt := promptui.Prompt{
			Label: "Enter Wiki ID",
		}
		inputID, err := prompt.Run()
		if err != nil {
			fmt.Println("Failed to read input")
			return
		}

		// konversi ID ke tipe int
		id, err := strconv.Atoi(inputID)
		if err != nil {
			fmt.Println("Invalid Wiki ID")
			return
		}

		// dapatkan wiki dari repository
		wiki, err := repo.GetWiki(id)
		if err != nil {
			fmt.Println("Failed to get wiki from database")
			return
		}
		createdAt, err := time.Parse("2006-01-02 15:04:05", wiki.CreatedAt)
		if err != nil {
			fmt.Println("Failed to parse created_at")
			return
		}

		updatedAt, err := time.Parse("2006-01-02 15:04:05", wiki.UpdatedAt)
		if err != nil {
			fmt.Println("Failed to parse updated_at")
			return
		}

		// tampilkan hasil pencarian ke pengguna
		fmt.Println("ID:", wiki.ID)
		fmt.Println("Topic:", wiki.Topic)
		fmt.Println("Description:", wiki.Description)
		fmt.Println("Created At:", createdAt)
		fmt.Println("Updated At:", updatedAt)
	}
}

func DeleteWikiHandler(repo repoWiki.WikiRepository) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		// Meminta pengguna untuk memasukkan ID wiki yang akan dihapus
		prompt := promptui.Prompt{
			Label: "Wiki ID",
		}
		idStr, err := prompt.Run()
		if err != nil {
			fmt.Println("Failed to read input")
			return
		}

		id, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Println("Invalid ID")
			return
		}

		// Menghapus wiki dari database
		err = repo.DeleteWiki(id)
		if err != nil {
			fmt.Println("Failed to delete data from database")
			return
		}

		// Menampilkan pesan sukses
		fmt.Println("Wiki dengan ID", id, "telah dihapus")
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func UpdateTopicDescHandler(repo repoWiki.WikiRepository) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		// Meminta pengguna untuk memasukkan ID topik yang akan diupdate
		prompt := promptui.Prompt{
			Label: "Enter the ID of the topic you want to update:",
		}
		idStr, err := prompt.Run()
		if err != nil {
			fmt.Println("Failed to read input")
			return
		}

		id, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Println("Invalid ID")
			return
		}

		// Meminta pengguna untuk memasukkan topik baru
		prompt = promptui.Prompt{
			Label: "Enter the new topic",
		}
		newTopic, err := prompt.Run()
		if err != nil {
			fmt.Println("Failed to read input")
			return
		}

		// Mengupdate topik
		err = repo.UpdateForWorker(id, newTopic)
		if err != nil {
			fmt.Println("Failed to update topic")
			return
		}

		res, err := http.Get("https://id.wikipedia.org/wiki/" + newTopic)
		if err != nil {
			fmt.Println("Gagal mengambil data dari Wikipedia")
			return
		}
		defer res.Body.Close()

		// Membuat dokumen dari respon HTTP
		doc, err := goquery.NewDocumentFromReader(res.Body)
		if err != nil {
			return
		}

		description := doc.Find("p").First().Text()

		// Mengupdate deskripsi dan kolom updated_at
		err = repo.UpdateDescriptionAndUpdatedAt(id, description)
		if err != nil {
			fmt.Println("Failed to update description and updated_at")
			return
		}

		fmt.Println("Topic and description have been updated")
	}
}

func WorkerHandler(repo repoWiki.WikiRepository) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		// ...

		// Query seluruh data dengan deskripsi kosong
		wikis, err := repo.GetWikisWithEmptyDescription()
		if err != nil {
			fmt.Println("Failed to get wikis")
			return
		}

		// Inisialisasi WaitGroup
		var wg sync.WaitGroup

		// Looping untuk setiap wiki dengan deskripsi kosong
		for _, wiki := range wikis {
			wg.Add(1) // Menambahkan jumlah goroutine yang akan dijalankan

			go func(wikiID int) {
				defer wg.Done() // Mengurangi jumlah goroutine yang sedang berjalan setelah selesai

				// Mengupdate deskripsi dari Wikipedia
				err := repo.UpdateDescriptionFromWikipedia(wikiID)
				if err != nil {
					fmt.Printf("Failed to update description for wiki ID %d\n", wikiID)
					return
				}

				// Mengupdate kolom updated_at
				err = repo.UpdateUpdatedAt(wikiID)
				if err != nil {
					fmt.Printf("Failed to update updated_at for wiki ID %d\n", wikiID)
				}
			}(wiki.ID)
		}

		// Menunggu semua goroutine selesai
		wg.Wait()

		// Cek apakah semua data sudah terisi

		fmt.Println("Worker finished")
	}
}
