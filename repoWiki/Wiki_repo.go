package repoWiki

import (
	"pr_ramadhan/cmd/models"
)

type WikiRepository interface {
	AddWiki(wiki *models.Wikis) error
	UpdateWiki(wiki *models.Wikis) error
	DeleteWiki(id uint) error
	GetWiki(id uint) (*models.Wikis, error)
}
