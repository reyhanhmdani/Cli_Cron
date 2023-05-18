package handlers

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/go-co-op/gocron"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"net/http"
	"net/url"
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

		// Menyimpan perubahan ke database
		err = repo.UpdateTopic(wiki)
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

func UpdateTopicHandler(repo repoWiki.WikiRepository) func(cmd *cobra.Command, args []string) {
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
		err = repo.UpdateTopic1(id, newTopic)
		if err != nil {
			fmt.Println("Failed to update topic")
			return
		}

		// Mengambil paragraf pertama dari Wikipedia
		doc, err := goquery.NewDocument("https://id.wikipedia.org/wiki/" + newTopic)
		if err != nil {
			fmt.Println("Failed to fetch data from Wikipedia")
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
		// Menjadwalkan tugas yang akan dijalankan setiap 1 menit
		s := gocron.NewScheduler(time.UTC)
		_, err := s.Every(1).Minute().Do(func() {
			// Mengambil data wiki dengan description kosong dari database
			wikis, err := repo.GetWikisWithEmptyDescription()
			if err != nil {
				fmt.Println("Failed to get data from database")
				return
			}

			// Concurrently melakukan http request dan update ke database
			var wg sync.WaitGroup
			for _, wiki := range wikis {
				wg.Add(1)
				go func(wiki *models.Wikis) {
					defer wg.Done()

					// Mengakses Wikipedia untuk mendapatkan paragraf pertama
					paragraph, err := fetchFirstParagraph(wiki.Topic)
					if err != nil {
						fmt.Println("Failed to fetch first paragraph:", err)
						return
					}

					// Update description dan updated_at di database
					wiki.Description = paragraph
					wiki.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
					err = repo.UpdateTopic(wiki)
					if err != nil {
						fmt.Println("Failed to update data in database")
						return
					}

					// Menampilkan pesan sukses
					fmt.Println("Description for topic", wiki.Topic, "has been updated")

					// cek jika semua deskripsi sudah terisi
					if isAllDescriptionsFilled(repo) {
						fmt.Println("Semua deskripsi sudah terisi")
						// menghentikan scheduler nya
						s.Stop()
					}
				}(wiki)
			}

			wg.Wait()
		})
		if err != nil {
			return
		}

		// Menjalankan scheduler
		s.StartBlocking()
	}
}

// Fungsi untuk mengambil paragraf pertama dari Wikipedia
func fetchFirstParagraph(topic string) (string, error) {
	nameUrl := "https://id.wikipedia.org/wiki/" + url.PathEscape(topic)
	resp, err := http.Get(nameUrl)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return "", err
	}

	paragraph := ""
	doc.Find("p").Each(func(i int, s *goquery.Selection) {
		if i == 0 {
			paragraph = s.Text()
			return
		}
	})

	return paragraph, nil
}

// Fungsi untuk memeriksa apakah semua deskripsi telah terisi
func isAllDescriptionsFilled(repo repoWiki.WikiRepository) bool {
	wikis, err := repo.GetWikisWithEmptyDescription()
	if err != nil {
		fmt.Println("Failed to get data from database")
		return false
	}
	return len(wikis) == 0
}
