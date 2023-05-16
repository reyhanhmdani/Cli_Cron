package database

import (
	"gorm.io/gorm"
	"pr_ramadhan/models"
	"pr_ramadhan/repoWiki"
)

type wikiRepository struct {
	db *gorm.DB
}

func NewWikiRepository(db *gorm.DB) repoWiki.WikiRepository {
	return &wikiRepository{db}
}

func (w *wikiRepository) AddWiki(wiki *models.Wiki) error {
	return w.db.Create(wiki).Error
}

func (w *wikiRepository) UpdateWiki(wiki *models.Wiki) error {
	return w.db.Save(wiki).Error
}

func (w *wikiRepository) DeleteWiki(id uint) error {
	return w.db.Delete(&models.Wiki{}, id).Error
}

func (w *wikiRepository) GetWiki(id uint) (*models.Wiki, error) {
	var wiki models.Wiki
	if err := w.db.First(&wiki, id).Error; err != nil {
		return nil, err
	}
	return &wiki, nil
}
