package handlers

import (
	"pr_ramadhan/models"
)

type WikiHandler struct {
	cfg *models.Config
}

func NewWikiHandlerImpl(cfg *models.Config) *WikiHandler {
	return &WikiHandler{
		cfg: cfg,
	}
}

func (h *WikiHandler) Create() error {
	//db, err := database.ConnnectDb(h.cfg)

	return nil
}
