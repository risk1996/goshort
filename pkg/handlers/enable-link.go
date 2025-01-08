package handlers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/risk1996/goshort/pkg/models"
	"github.com/risk1996/goshort/pkg/utils"
	"gorm.io/gorm/clause"
)

func EnableLink(c *gin.Context) {
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

	db.
		Clauses(clause.Returning{}).
		Model(&entry).
		Unscoped().
		Where("path = ? AND admin_secret = ?", path, req.Secret).
		Update("deleted_at", sql.NullTime{Valid: false})
	c.JSON(http.StatusOK, MapToResponse(&entry))
}
