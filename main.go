package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// album is the data about a record album

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// main function
// Gin router is initialized using `Default`
// Get function is used to associate GET HTTP method and /albums path with a handler function
// `getAlbums` is the function being passed while `getAlbums()` is the result of that funcitons
// Run function is used to attach the router to an http.Server and start the server

func main() {

	router := gin.Default()
	// Gin allows associating the handler with an HTTP method and path combination
	// This way you can separately route requests sent to a single path based on the method being used by the client
	router.GET("/albums", getAlbums)
	// associate :id for handler
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)
	router.Run("localhost:8080")

}

var albums = []album{
	{ID: "1", Title: "Revolver", Artist: "The Beatles", Price: 56.99},
	{ID: "2", Title: "Low", Artist: "David Bowie", Price: 34.99},
	{ID: "3", Title: "Exile on Main St", Artist: "The Rolling Stones", Price: 76.99},
}

// getAlbums responds with list of all albums in JSON
// function takes gin.Context parameter
// gin.Context carries request details, validates, serializes JSON etc. (different from context package in go)
// Context.IndentedJSON is called to serialize the struct into JSON and add it to response
// HTTP status code 200 OK is sent
// replace `IndentedJSON` to `JSON` to send more compact JSON; though size difference is small but readability diff is big

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func postAlbums(c *gin.Context) {
	var newAlbum album
	// Call bindJSON to bind the received JSON to newalbum
	// This binds the request body to newAlbum
	// StatusCreated is a 201 status code along with the JSON
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)

}

func getAlbumByID(c *gin.Context) {
	// Context.Param retrieves the id path from the url; when we map this handler to a path, we include the placeholder for the parameter in path
	id := c.Param("id")
	// Loop over albums looking for matching id parameter value
	// Serialize that album struct to JSON and return with 200 OK response
	// In the real world, a daabase query would be used to perform this lookup
	// 404 Error with http.StatusNotFound if album isn't found
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Album not found!"})
}
