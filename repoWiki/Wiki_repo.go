package repoWiki

import "pr_ramadhan/models"

type WikiRepository interface {
	AddWiki(wiki *models.Wiki) error
	UpdateWiki(wiki *models.Wiki) error
	DeleteWiki(id uint) error
	GetWiki(id uint) (*models.Wiki, error)
}
