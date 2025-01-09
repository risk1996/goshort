package main

import (
	"flag"
	"fmt"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/risk1996/goshort/pkg/controller"
	"github.com/risk1996/goshort/pkg/docs"
	"github.com/risk1996/goshort/pkg/utils"
)

// @Title						Goshort
// @Version					0.1
// @Description				A simple Go-based service to manage shortened links with support for create, access, edit, enable, and disable functionalities.
// @Contact.Name				William Darian
// @Contact.Email				williamdariansutedjo@gmail.com
// @BasePath					/
// @ExternalDocs.Description	GitHub
// @ExternalDocs.URL			https://github.com/risk1996/goshort
func main() {
	port := flag.Int("port", 8080, "Server port")
	flag.Parse()

	db := utils.ConnectAndMigrateDatabase("db.db")
	r := gin.Default()
	c := controller.NewController()
	utils.AttachDB(r, db)

	docs.SwaggerInfo.Host = fmt.Sprintf("localhost:%v", *port)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.PUT("/", c.PutLink)
	r.GET("/:path", c.AccessLink)
	r.PATCH("/:path/edit", c.EditLink)
	r.PATCH("/:path/disable", c.DisableLink)
	r.PATCH("/:path/enable", c.EnableLink)

	r.Run(fmt.Sprintf(":%v", *port))
}
