package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/risk1996/goshort/pkg/core"
	"github.com/risk1996/goshort/pkg/models"
	"github.com/risk1996/goshort/pkg/utils"
	"gorm.io/gorm/clause"
)

// EditLink godoc
//
//	@Summary		Edit a link
//	@Description	Edits the target URL for the given path and secret. Idempotent.
//	@Tags			link
//	@Param			path			path	string	true	"Shortened path"
//	@Param			Authorization	header	string	true	"Bearer token containing admin secret"
//	@Accept			json
//	@Param			body	body	ShortenLinkRequest	true	"Target URL"
//	@Produce		json
//	@Success		200	{object}	LinkResponse	"Link edited successfully."
//	@Failure		400	"Invalid request."
//	@Failure		401	"Invalid admin secret."
//	@Failure		404	"Link not found."
//	@Router			/{path}/edit [patch]
func (*Controller) EditLink(c *gin.Context) {
	db := utils.GetDB(c)
	path := c.Params.ByName("path")
	secret, err := utils.ParseAuthBearerToken(c)

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	var req ShortenLinkRequest
	if c.Bind(&req) != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	url, err := core.NormalizeURL(req.URL)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	entry := models.Link{}
	res := db.
		Limit(1).
		Where("path = ?", path).
		Find(&entry)
	if res.RowsAffected == 0 {
		c.AbortWithStatus(http.StatusNotFound)
		return
	} else if entry.AdminSecret != secret {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	db.
		Model(&entry).
		Clauses(clause.Returning{}).
		Where("path = ? AND admin_secret = ?", path, secret).
		Update("target", url)
	c.JSON(http.StatusOK, MapToResponse(&entry))
}
