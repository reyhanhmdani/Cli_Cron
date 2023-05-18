package repoWiki

import (
	"pr_ramadhan/cmd/models"
)

type WikiRepository interface {
	AddWiki(wiki *models.Wikis) error
	UpdateTopic(wiki *models.Wikis) error
	DeleteWiki(id int) error
	GetWiki(id int) (*models.Wikis, error)
	GetWikisWithEmptyDescription() ([]*models.Wikis, error)
	UpdateDescriptionByTopic(topic, description string) error
	UpdateTopic1(id int, newTopic string) error
	UpdateDescriptionAndUpdatedAt(id int, description string) error
}
