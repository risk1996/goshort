package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/risk1996/goshort/pkg/core"
	"github.com/risk1996/goshort/pkg/models"
	"github.com/risk1996/goshort/pkg/utils"
)

type Request struct {
	URL string `json:"url" binding:"required,http_url"`
}

func PutLink(c *gin.Context) {
	db := utils.GetDB(c)

	var req Request
	if c.BindJSON(&req) == nil {
		url, err := core.NormalizeURL(req.URL)
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		path := core.RandShortLinkPath(8)
		secret := uuid.New().String()
		entry := models.Link{Path: path, Target: url, AdminSecret: secret}
		db.Create(&entry)

		c.JSON(http.StatusOK, entry)
	} else {
		c.AbortWithStatus(http.StatusBadRequest)
	}
}
