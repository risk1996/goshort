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
//	@Param			path	path	string	true	"Shortened path"
//	@Accept			json
//	@Param			body	body	EditLinkRequest	true	"Secret and new target URL"
//	@Produce		json
//	@Success		200	{object}	LinkResponse	"Link edited successfully."
//	@Failure		400	"Invalid request."
//	@Failure		404 "Link not found or wrong secret."
//	@Router			/{path}/edit [patch]
func (*Controller) EditLink(c *gin.Context) {
	db := utils.GetDB(c)
	path := c.Params.ByName("path")

	var req EditLinkRequest
	if c.Bind(&req) != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	url, err := core.NormalizeURL(req.URL)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	var entry models.Link
	res := db.
		Model(&entry).
		Clauses(clause.Returning{}).
		Where("path = ? AND admin_secret = ?", path, req.Secret).
		Update("target", url)
	if res.RowsAffected == 0 {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, MapToResponse(&entry))
}
