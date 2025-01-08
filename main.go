package main

import (
	"github.com/gin-gonic/gin"

	"github.com/risk1996/goshort/pkg/handlers"
	"github.com/risk1996/goshort/pkg/utils"
)

func main() {
	db := utils.ConnectAndMigrateDatabase("db.db")
	r := gin.Default()
	utils.AttachDB(r, db)

	r.PUT("/", handlers.PutLink)
	r.GET("/:path", handlers.AccessLink)

	r.Run(":8080")
}
