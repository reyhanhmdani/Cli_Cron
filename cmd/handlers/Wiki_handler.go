package handlers

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"net/http"
	"pr_ramadhan/cmd/models"
	"pr_ramadhan/repoWiki"
	"strconv"
)

// struct ini untuk melakukan scarping paragraf pertama dari Wiki
type Scrapper struct {
	Client *http.Client
}

// //////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
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
		//now := time.Now()
		wiki := models.Wikis{
			Topic: topic,
			//CreatedAt: now.Format(time.DateTime),
			//UpdatedAt: now.Format(time.DateTime),
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
