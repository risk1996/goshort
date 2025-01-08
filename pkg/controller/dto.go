package controller

import "github.com/risk1996/goshort/pkg/models"

type ShortenLinkRequest struct {
	// The target URL to shorten.
	URL string `json:"url" binding:"required,http_url" validate:"required" format:"url" example:"https://www.google.com"`
} //	@name ShortenLinkRequest

type LinkResponse struct {
	// The shortened path.
	Path string `json:"path" validate:"required" example:"abc12345"`
	// The target URL.
	Target string `json:"target" validate:"required" format:"url" example:"https://www.google.com"`
	// Whether the link is active.
	Active bool `json:"active" validate:"required" example:"true"`
	// The secret key for managing the link.
	Secret string `json:"secret" validate:"required" format:"uuid" example:"bcc9a044-918a-4ffa-ae4b-75274ca23668"`
} //	@name LinkResponse

func MapToResponse(m *models.Link) LinkResponse {
	return LinkResponse{
		Path:   m.Path,
		Target: m.Target,
		Active: !m.DeletedAt.Valid,
		Secret: m.AdminSecret,
	}
}

type AdminRequest struct {
	// The secret key for managing the link.
	Secret string `json:"secret" validate:"required" format:"uuid" example:"bcc9a044-918a-4ffa-ae4b-75274ca23668"`
} //	@name AdminRequest

type EditLinkRequest struct {
	AdminRequest
	ShortenLinkRequest
} //	@name EditLinkRequest
