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
//	@Param			path			path	string	true	"Shortened path"
//	@Param			Authorization	header	string	true	"Bearer token containing admin secret"
//	@Produce		json
//	@Success		200	{object}	LinkResponse	"Link disabled successfully."
//	@Failure		401	"Invalid admin secret."
//	@Failure		404	"Link not found."
//	@Router			/{path}/disable [patch]
func (*Controller) DisableLink(c *gin.Context) {
	db := utils.GetDB(c)
	path := c.Params.ByName("path")
	secret, err := utils.ParseAuthBearerToken(c)

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	entry := models.Link{}
	res := db.
		Limit(1).
		Unscoped().
		Where("path = ?", path).
		Find(&entry)
	if res.RowsAffected == 0 {
		c.AbortWithStatus(http.StatusNotFound)
		return
	} else if entry.AdminSecret != secret {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	db.Delete(&entry)
	c.JSON(http.StatusOK, MapToResponse(&entry))
}
