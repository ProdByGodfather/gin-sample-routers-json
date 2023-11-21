package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type post struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

var posts = []post{
	{ID: "1", Title: "Whats the best Movies", Body: "The Dark knight"},
	{ID: "2", Title: "Hello World", Body: "The First Json"},
	{ID: "3", Title: "AbarVision", Body: "Abarvision for the Godfather"},
}

func main() {
	router := gin.Default()

	router.GET("posts", getAllPosts)
	router.GET("posts/:id", getPostByID)
	err := router.Run("localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
}

func getAllPosts(c *gin.Context) {
	/*
		Intended JSON serializers the given struct as pretty JSON (indented + endlines)
		Ingo the response body. it also sets the context-type as "application/json".
		WARNING: we recommend to use this only for development purposes since printing.
		pretty JSON is more CPU and bandwidth consuming. use context.JSON() instead.
		func (c *Context) IntendedJSON(code int,obj interface{})
	*/
	c.IndentedJSON(http.StatusOK, posts)
}
func getPostByID(c *gin.Context) {
	id := c.Param("id")
	// se arch for particular posts in slice for demo (db in production)
	var index int
	for k, v := range posts {
		if id == v.ID {
			index = k
		}
	}
	c.IndentedJSON(http.StatusOK, posts[index])
}
