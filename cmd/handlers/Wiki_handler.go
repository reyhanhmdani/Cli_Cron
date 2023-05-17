package handlers

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"net/http"
	"pr_ramadhan/cmd/models"
	"pr_ramadhan/repoWiki"
	"strconv"
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

//func ScrapeDescription(db *gorm.DB) {
//	// Query seluruh data dengan deskripsi kosong
//	var wikis []models.Wikis
//	db.Where("description IS NULL").Find(&wikis)
//
//	// Buat HTTP client
//	client := http.DefaultClient
//
//	// Lakukan scraping dan update deskripsi secara koncurrent
//	var wg sync.WaitGroup
//	for _, wiki := range wikis {
//		wg.Add(1)
//		go func(wiki models.Wikis) {
//			defer wg.Done()
//
//			// Buat URL Wikipedia berdasarkan topik
//			url := fmt.Sprintf("https://id.wikipedia.org/wiki/%s", wiki.Topic)
//
//			// Lakukan request HTTP
//			resp, err := client.Get(url)
//			if err != nil {
//				fmt.Printf("Failed to scrape %s: %s\n", wiki.Topic, err.Error())
//				return
//			}
//			defer func(Body io.ReadCloser) {
//				err := Body.Close()
//				if err != nil {
//
//				}
//			}(resp.Body)
//
//			// Baca body response menggunakan goquery
//			doc, err := goquery.NewDocumentFromReader(resp.Body)
//			if err != nil {
//				fmt.Printf("Failed to parse response body for %s: %s\n", wiki.Topic, err.Error())
//				return
//			}
//
//			// Ambil paragraf pertama dari body
//			paragraph := doc.Find("p").First().Text()
//
//			// Update deskripsi dan kolom updated_at pada tabel wikis
//			wiki.Description = paragraph
//			now := time.Now()
//			wiki.UpdatedAt = now.Format(time.DateTime)
//			if err := db.Save(&wiki).Error; err != nil {
//				fmt.Printf("Failed to update description for %s: %s\n", wiki.Topic, err.Error())
//				return
//			}
//
//			fmt.Printf("Description updated for %s\n", wiki.Topic)
//		}(wiki)
//	}
//
//	// Tunggu hingga semua scraping dan update selesai
//	wg.Wait()
//}

//func StartWorker(db *gorm.DB) {
//	// Buat scheduler dengan interval 1 menit
//	s := gocron.NewScheduler(time.UTC)
//	_, err := s.Every(1).Minute().Do(func() {
//		ScrapeDescription(db)
//	})
//	if err != nil {
//		fmt.Println("Failed to schedule worker:", err.Error())
//		return
//	}
//
//	// Jalankan scheduler
//	s.StartAsync()
//	fmt.Println("Worker started")
//}
