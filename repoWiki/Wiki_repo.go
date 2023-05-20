package repoWiki

import (
	"pr_ramadhan/cmd/models"
)

type WikiRepository interface {
	GetAllWikis() ([]*models.Wikis, error)
	AddWiki(wiki *models.Wikis) error
	UpdateWiki(wiki *models.Wikis) error
	DeleteWiki(id int) error
	GetWiki(id int) (*models.Wikis, error)
	GetWikisWithEmptyDescription() ([]*models.Wikis, error)
	//UpdateDescriptionByTopic(topic, description string) error
	UpdateForWorker(id int, newTopic string) error
	UpdateDescriptionAndUpdatedAt(id int, description string) error
	UpdateDescriptionFromWikipedia(id int) error

	// ONLY WORKER
	UpdateUpdatedAt(id int) error
}
