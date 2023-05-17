package database

import (
	"gorm.io/gorm"
	"pr_ramadhan/models"
	"pr_ramadhan/repoWiki"
	"time"
)

type wikiRepository struct {
	db *gorm.DB
}

func NewWikiRepository(db *gorm.DB) repoWiki.WikiRepository {
	return &wikiRepository{db}
}

func (w *wikiRepository) AddWiki(wiki *models.Wikis) error {
	return w.db.Create(wiki).Error
}

func (w *wikiRepository) UpdateWiki(wiki *models.Wikis) error {
	return w.db.Save(wiki).Error
}

func (w *wikiRepository) DeleteWiki(id uint) error {
	return w.db.Delete(&models.Wikis{}, id).Error
}

func (w *wikiRepository) GetWiki(id uint) (*models.Wikis, error) {
	var wiki models.Wikis
	if err := w.db.First(&wiki, id).Error; err != nil {
		return nil, err
	}

	createdAt, err := time.Parse(time.RFC3339, wiki.CreatedAt)
	if err != nil {
		return nil, err
	}

	wiki.CreatedAt = createdAt.Format(time.RFC3339)
	return &wiki, nil
}
