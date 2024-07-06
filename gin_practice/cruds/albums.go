package cruds

import (
	"github.com/BinDruid/go-practice/gin_practice/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type Album struct {
	gorm.Model
	Title  string `json:"title"`
	Artist string `json:"artist"`
	Price  uint   `json:"price"`
}

func GetAllAlbums(c *gin.Context) {
	var albums []models.Album
	models.DB.Find(&albums)
	c.IndentedJSON(http.StatusOK, albums)
}

func CreateAlbum(c *gin.Context) {
	var newAlbum Album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}
	album := models.Album{Title: newAlbum.Title, Artist: newAlbum.Artist, Price: newAlbum.Price}
	models.DB.Create(&album)
	c.IndentedJSON(http.StatusCreated, album)
}

func GetAlbumByID(c *gin.Context) {
	var album models.Album
	if err := models.DB.Where("ID = ?", c.Param("id")).First(&album).Error; err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": album})
}
