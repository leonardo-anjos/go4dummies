package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// album represents data about a record album.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albumsMap = make(map[string]album)

// albums slice to seed record album data.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func isDuplicate(title, artist string) bool {
	for _, existingAlbum := range albumsMap {
		if existingAlbum.Title == title && existingAlbum.Artist == artist {
			return true
		}
	}
	return false
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)
	router.DELETE("/albums/:id", deleteAlbumByID)

	router.Run(":8080")
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
	var newAlbum album

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid request payload"})
		return
	}

	// Validate uniqueness based on title and artist
	if isDuplicate(newAlbum.Title, newAlbum.Artist) {
		c.IndentedJSON(http.StatusConflict, gin.H{"error": "album with the same title and artist already exists"})
		return
	}

	// Add the new album to the slice.
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func deleteAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// Find the index of the album with the given ID
	index := -1
	for i, album := range albums {
		if album.ID == id {
			index = i
			break
		}
	}

	// If the album is found, remove it from the slice
	if index != -1 {
		albums = append(albums[:index], albums[index+1:]...)
		c.IndentedJSON(http.StatusOK, gin.H{"message": fmt.Sprintf("album with ID %s deleted", id)})
	} else {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("album with ID %s not found", id)})
	}
}
