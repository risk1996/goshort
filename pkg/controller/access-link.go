package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/risk1996/goshort/pkg/models"
	"github.com/risk1996/goshort/pkg/utils"
)

// AccessLink godoc
//
//	@Summary		Access a link
//	@Description	Redirects to the original URL for the given shortened path.
//	@Tags			link
//	@Param			path	path	string	true	"Shortened path"
//	@Success		301		"Redirects to the target URL."
//	@Failure		404		"Link not found or inactive."
//	@Router			/{path} [get]
func (*Controller) AccessLink(c *gin.Context) {
	db := utils.GetDB(c)
	path := c.Params.ByName("path")

	var entry models.Link
	res := db.Limit(1).Find(&entry, "deleted_at IS NULL AND path = ?", path)
	if res.RowsAffected == 0 {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.Redirect(http.StatusMovedPermanently, entry.Target)
	}
}
