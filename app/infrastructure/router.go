package infrastructure

import (
	"go_http_api_template/interfaces/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func init() {
	router := gin.Default()
	router.Use(cors.Default())

	userController := controllers.NewUserController(NewSqlHandler())

	router.GET("/", func(c *gin.Context) {
		c.String(200, "Hello,World!")
	})
	// ルーティング
	u := router.Group("/users")
	{
		u.GET("/", func(c *gin.Context) { userController.GetUser() })
	}

	Router = router
}
