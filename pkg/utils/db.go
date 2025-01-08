package utils

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/risk1996/goshort/pkg/models"
)

func ConnectAndMigrateDatabase(name string) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(name), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database.")
	}

	err = db.AutoMigrate(models.AllModelReferences()...)
	if err != nil {
		panic("Failed to migrate database to the latest state.")
	}

	return db
}

const DB_CONTEXT_KEY = "DB"

func AttachDB(e *gin.Engine, db *gorm.DB) {
	e.Use(func(c *gin.Context) {
		c.Set(DB_CONTEXT_KEY, db)
		c.Next()
	})
}

func GetDB(c *gin.Context) *gorm.DB {
	return c.MustGet(DB_CONTEXT_KEY).(*gorm.DB)
}
