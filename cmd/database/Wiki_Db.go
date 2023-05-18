package database

import (
	"gorm.io/gorm"
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

func (w *wikiRepository) AddWiki(wiki *models.Wikis) error {
	return w.db.Create(wiki).Error
}

func (w *wikiRepository) UpdateTopic(wiki *models.Wikis) error {
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

func (w *wikiRepository) UpdateDescriptionByTopic(topic, description string) error {
	return w.db.Model(&models.Wikis{}).Where("topic = ?", topic).Update("description", description).Error
}
func (w *wikiRepository) UpdateTopic1(id int, newTopic string) error {
	return w.db.Model(&models.Wikis{}).Where("id = ?", id).Update("topic", newTopic).Error
}

func (w *wikiRepository) UpdateDescriptionAndUpdatedAt(id int, description string) error {
	now := time.Now()

	return w.db.Model(&models.Wikis{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"description": description,
			"updated_at":  now,
		}).Error
}
