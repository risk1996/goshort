package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/risk1996/goshort/pkg/core"
	"github.com/risk1996/goshort/pkg/models"
	"github.com/risk1996/goshort/pkg/utils"
)

// PutLink godoc
//
//	@Summary		Create a new link
//	@Description	Creates a new shortened link for the provided URL.
//	@Tags			link
//	@Accept			json
//	@Param			body	body	ShortenLinkRequest	true	"Target URL"
//	@Produce		json
//	@Success		200	{object}	LinkResponse	"Link created successfully."
//	@Failure		400	"Invalid request."
//	@Router			/ [put]
func (*Controller) PutLink(c *gin.Context) {
	db := utils.GetDB(c)

	var req ShortenLinkRequest
	if c.BindJSON(&req) != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	url, err := core.NormalizeURL(req.URL)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	path := core.RandShortLinkPath(8)
	secret := uuid.New().String()
	entry := models.Link{Path: path, Target: url, AdminSecret: secret}
	db.Create(&entry)

	c.JSON(http.StatusOK, MapToResponse(&entry))
}
