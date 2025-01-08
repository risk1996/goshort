package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/risk1996/goshort/pkg/core"
	"github.com/risk1996/goshort/pkg/models"
	"github.com/risk1996/goshort/pkg/utils"
	"gorm.io/gorm/clause"
)

func EditLink(c *gin.Context) {
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
