package router

import (
	"vinyl-api/controllers"

	"github.com/gin-gonic/gin"
)

var Router = gin.Default()

func SetRoutes() {
	Router.GET("/albums", controllers.GetAlbums)
	Router.GET("/albums/:id", controllers.GetAlbumByID)
	Router.POST("/albums", controllers.AddAlbum)
}

func RunRouter() {
	Router.Run("localhost:8080")

}
