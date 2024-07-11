package routes

import (
	"github.com/BinDruid/go-practice/cruds"
	"github.com/gin-gonic/gin"
)

func AddAlbumsRoute(rg *gin.RouterGroup, path string) {
	albumsRouter := rg.Group(path)
	albumsRouter.GET("/", cruds.GetAllAlbums)
	albumsRouter.GET("/:id/", cruds.GetAlbumByID)
	albumsRouter.POST("/", cruds.CreateAlbum)
}
