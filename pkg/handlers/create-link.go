package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/risk1996/goshort/pkg/core"
	"github.com/risk1996/goshort/pkg/models"
	"github.com/risk1996/goshort/pkg/utils"
)

type Request struct {
	URL string `json:"url" binding:"required,http_url"`
}

func CreateLink(c *gin.Context) {
	db := utils.GetDB(c)

	var req Request
	if c.BindJSON(&req) == nil {
		url, err := core.NormalizeURL(req.URL)
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		entry := models.Link{Path: core.RandShortLinkPath(8), Target: url}
		db.FirstOrCreate(&entry, "target = ?", url)

		c.JSON(http.StatusCreated, entry)
	} else {
		c.AbortWithStatus(http.StatusBadRequest)
	}
}
