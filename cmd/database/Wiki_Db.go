package database

import "gorm.io/gorm"

import (
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"pr_ramadhan/cmd/models"
	"pr_ramadhan/repoWiki"
	"time"
)

type wikiRepository struct {
	db *gorm.DB
}

func NewWikiRepository(db *gorm.DB) repoWiki.WikiRepository {
	return &wikiRepository{db}
}

//func (w *wikiRepository) GetAllWikis() ([]*models.Wikis, error) {
//	var wikis []*models.Wikis
//	err := w.db.Find(&wikis).Error
//	if err != nil {
//		return nil, err
//	}
//	return wikis, nil
//}

func (w *wikiRepository) AddWiki(wiki *models.Wikis) error {
	return w.db.Create(wiki).Error
}

func (w *wikiRepository) UpdateWiki(wiki *models.Wikis) error {
	return w.db.Save(wiki).Error
}

func (w *wikiRepository) DeleteWiki(id int) error {
	return w.db.Delete(&models.Wikis{}, id).Error
}

func (w *wikiRepository) GetWiki(id int) (*models.Wikis, error) {
	wiki := &models.Wikis{}
	err := w.db.First(wiki, id).Error
	if err != nil {
		return nil, err
	}
	return wiki, nil
}

func (w *wikiRepository) GetWikisWithEmptyDescription() ([]*models.Wikis, error) {
	var wikis []*models.Wikis
	err := w.db.Where("description IS NULL OR description = ?", "").Find(&wikis).Error
	if err != nil {
		return nil, err
	}
	return wikis, nil
}

//	func (w *wikiRepository) UpdateDescriptionByTopic(topic, description string) error {
//		return w.db.Model(&models.Wikis{}).Where("topic = ?", topic).Update("description", description).Error
//	}
func (w *wikiRepository) UpdateForWorker(id int, newTopic string) error {
	return w.db.Model(&models.Wikis{}).Where("id = ?", id).Update("topic", newTopic).Error
}

func (w *wikiRepository) UpdateDescriptionAndUpdatedAt(id int, description string) error {
	currentTime := time.Now().UTC()

	return w.db.Model(&models.Wikis{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"description": description,
			"updated_at":  currentTime,
		}).Error
}

func (w *wikiRepository) UpdateDescriptionFromWikipedia(id int) error {
	wiki, err := w.GetWiki(id)
	if err != nil {
		return err
	}

	// Mengambil halaman Wikipedia
	res, err := http.Get("https://id.wikipedia.org/wiki/" + wiki.Topic)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	// Membuat dokumen dari respon HTTP
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return err
	}

	description := doc.Find("p").First().Text()

	// Memperbarui deskripsi jika berbeda
	if wiki.Description != description || wiki.Description == "" {
		wiki.Description = description
		err := w.db.Save(wiki).Error
		if err != nil {
			return err
		}
	}

	return nil
}

// WORKER
func (w *wikiRepository) UpdateUpdatedAt(id int) error {
	currentTime := time.Now().UTC()

	err := w.db.Model(&models.Wikis{}).Where("id = ?", id).Update("updated_at", currentTime).Error
	if err != nil {
		return err
	}

	return nil
}
