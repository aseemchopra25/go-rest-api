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
	router.GET("/albums", getAlbums)
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
