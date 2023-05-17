package repoWiki

import (
	"pr_ramadhan/cmd/models"
)

type WikiRepository interface {
	AddWiki(wiki *models.Wikis) error
	UpdateWiki(wiki *models.Wikis) error
	DeleteWiki(id int) error
	GetWiki(id int) (*models.Wikis, error)
}
