package database

import (
	"gorm.io/gorm"
	"pr_ramadhan/cmd/models"
	"pr_ramadhan/repoWiki"
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
	return &wiki, nil
}
