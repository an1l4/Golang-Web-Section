package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//album represents data about a record album

type Album struct {
	AlbumID string  `json:"id"`
	Title   string  `json:"title"`
	Artist  string  `json:"artist"`
	Price   float64 `json:"price"`
}

//album slices to seed record album data

var albums = []Album{
	{AlbumID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{AlbumID: "2", Title: "Jeru", Artist: "Gerry", Price: 17.99},
	{AlbumID: "3", Title: "Sarah and Clifford", Artist: "Sarah", Price: 39.99},
}

//getAlbum records with the list of all albums as json

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)

}

//implementing 1st endpoint

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)
	router.Run("localhost:8080")

}

//postAlbums adds an album from json received in the request body
func postAlbums(c *gin.Context) {
	var newAlbum Album

	//call BindJson to bind the received json to newAlbum

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	//add new album to the slice
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

//getAlbumById locates the album whose ID value matches the id
//parameter sent by the client,then return  that album  as a response
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	//loop over the  list of  albums ,looking for
	//an album whose ID  value matches the parameter

	for _, a := range albums {
		if a.AlbumID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})

}
