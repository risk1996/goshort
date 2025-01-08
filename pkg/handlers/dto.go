package handlers

import "github.com/risk1996/goshort/pkg/models"

type PutLinkRequest struct {
	URL string `json:"url" binding:"required,http_url"`
}

type LinkResponse struct {
	Path   string `json:"path"`
	Target string `json:"target"`
	Active bool   `json:"active"`
	Secret string `json:"secret"`
}

func MapToResponse(m *models.Link) LinkResponse {
	return LinkResponse{
		Path:   m.Path,
		Target: m.Target,
		Active: !m.DeletedAt.Valid,
		Secret: m.AdminSecret,
	}
}

type AdminRequest struct {
	Secret string `json:"secret"`
}
