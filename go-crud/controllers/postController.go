package controllers

import (
	"go-crud/initializers"
	"go-crud/models"
	"log"

	"github.com/gin-gonic/gin"
)

func HandlePingRequest(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func CreatePost(c *gin.Context) {

	var newPost models.Post
	err := c.BindJSON(&newPost)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	log.Print("Creating new post with title: ", newPost.Title)

	result := initializers.DB.Create(&newPost)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(201, gin.H{
		"status":  "success",
		"message": "New post created",
		"data":    newPost,
	})
}

func GetPost(c *gin.Context) {
	id := c.Param("id")

	var post models.Post

	result := initializers.DB.First(&post, id)
	if result.Error != nil {
		c.JSON(404, gin.H{
			"status":  "error",
			"message": "Post not found",
		})
		return
	}

	c.JSON(200, gin.H{
		"status": "success",
		"data":   post,
	})
}
