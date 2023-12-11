package controllers

import (
	"net/http"
	"strconv"
	"vinyl-api/structs"

	"github.com/gin-gonic/gin"
)

func GetAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, structs.Albums)
}

func GetAlbumByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "Bad ID provided",
		})
		return
	}
	for _, album := range structs.Albums {
		if album.Id == id {
			c.IndentedJSON(http.StatusOK, album)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{
		"message": " Album not found",
	})
}

func AddAlbum(c *gin.Context) {
	var newAlbum structs.Album
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}
	structs.Albums = append(structs.Albums, newAlbum)

	c.IndentedJSON(http.StatusCreated, newAlbum)
}
