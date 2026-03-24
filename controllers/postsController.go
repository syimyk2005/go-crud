package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rubbyklein/go-crud/initializers"
	"github.com/rubbyklein/go-crud/models"
	"gorm.io/gorm"
)

func PostsCreate(c *gin.Context) {

	var body struct {
		Title string `json:"title"`
		Body  string `json:"body"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	post := models.Post{
		Title: body.Title,
		Body:  body.Body,
	}

	if err := initializers.DB.Create(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create post"})
		return
	}

	c.JSON(201, gin.H{
		"post": post,
	})
}

func PostsIndex(c *gin.Context) {

	var posts []models.Post
	if err := initializers.DB.Find(&posts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch posts"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"posts": posts,
	})
}

func PostsById(c *gin.Context) {
	id := c.Param("id")

	var post models.Post
	if err := initializers.DB.First(&post, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "post not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch post"})
		return
	}

	c.JSON(http.StatusOK, post)
}

func PostsUpdate(c *gin.Context) {
	id := c.Param("id")

	var body struct {
		Title string `json:"title"`
		Body  string `json:"body"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var post models.Post

	if err := initializers.DB.First(&post, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "post with id " + id + " does not exist"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch post"})
		return
	}

	post.Title = body.Title
	post.Body = body.Body

	if err := initializers.DB.Save(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update post"})
		return
	}

	c.JSON(http.StatusOK, post)
}

func PostsDelete(c *gin.Context) {
	id := c.Param("id")

	var post models.Post
	if err := initializers.DB.First(&post, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "post with id " + id + " does not exist or already was deleted"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to find post"})
		return
	}
	
	if err := initializers.DB.Delete(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "post was not deleted"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your post with id " + id + " was deleted",
	})
}
