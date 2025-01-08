package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/risk1996/goshort/pkg/models"
	"github.com/risk1996/goshort/pkg/utils"
)

// DisableLink godoc
//
//	@Summary		Disable a link
//	@Description	Disables the link with the given path and secret. Idempotent.
//	@Tags			link
//	@Param			path	path	string	true	"Shortened path"
//	@Accept			json
//	@Param			body	body	AdminRequest	true	"Secret"
//	@Produce		json
//	@Success		200	{object}	LinkResponse	"Link disabled successfully."
//	@Failure		400	"Invalid request."
//	@Failure		404 "Link not found or wrong secret."
//	@Router			/{path}/disable [patch]
func (*Controller) DisableLink(c *gin.Context) {
	db := utils.GetDB(c)
	path := c.Params.ByName("path")

	var req AdminRequest
	if c.Bind(&req) != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	entry := models.Link{}
	res := db.
		Limit(1).
		Unscoped().
		Where("path = ? AND admin_secret = ?", path, req.Secret).
		Find(&entry)
	if res.RowsAffected == 0 {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	db.Delete(&entry)
	c.JSON(http.StatusOK, MapToResponse(&entry))
}
