package main

import (
	"fmt"
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
	router.POST("posts", createPost)
	router.PATCH("posts/:id", updatePost)
	router.DELETE("posts/:id", deletePost)

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
func createPost(c *gin.Context) {
	var newPost post
	err := c.BindJSON(&newPost)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	posts = append(posts, newPost)
	c.IndentedJSON(http.StatusCreated, posts)
}

func updatePost(c *gin.Context) {
	var index int
	var update post
	err := c.BindJSON(&update)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	id := c.Param("id")
	for k, v := range posts {
		if id == v.ID {
			index = k
		}
	}

	title := update.Title
	body := update.Body
	// update item
	posts[index].Title = title
	posts[index].Body = body
	c.IndentedJSON(http.StatusOK, posts[index])
}
func deletePost(c *gin.Context) {
	var index int
	id := c.Param("id")
	for k, v := range posts {
		if id == v.ID {
			index = k
		}
	}
	// delete item from slice
	posts = append(posts[:index], posts[index+1:]...)
	c.IndentedJSON(http.StatusOK, posts)
}
