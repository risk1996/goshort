package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/risk1996/goshort/pkg/core"
	"github.com/risk1996/goshort/pkg/models"
	"github.com/risk1996/goshort/pkg/utils"
)

func PutLink(c *gin.Context) {
	db := utils.GetDB(c)

	var req PutLinkRequest
	if c.BindJSON(&req) != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	url, err := core.NormalizeURL(req.URL)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	path := core.RandShortLinkPath(8)
	secret := uuid.New().String()
	entry := models.Link{Path: path, Target: url, AdminSecret: secret}
	db.Create(&entry)

	c.JSON(http.StatusOK, MapToResponse(&entry))
}
