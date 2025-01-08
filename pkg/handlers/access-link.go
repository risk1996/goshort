package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/risk1996/goshort/pkg/models"
	"github.com/risk1996/goshort/pkg/utils"
)

func AccessLink(c *gin.Context) {
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
