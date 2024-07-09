package controllers

import (
	"dunky-star/go-simple/initializers"
	"dunky-star/go-simple/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type createPostRequest struct {
	Title string `json:"title" binding:"required"`
	Body string `json:"body" binding:"required"`
}

// type postResponse struct {
// 	Title string `json:"title"`
// 	Body string  `json:"body"`
// }


func PostCreate(ctx *gin.Context) {
	// Get data off request body
    var req createPostRequest

	if err := ctx.ShouldBindJSON(&req); err != nil{
		ctx.JSON(http.StatusBadRequest, err)
        return
	}
	// Create a post
	posts := models.Post{Title: req.Title, Body: req.Body}
    
	result := initializers.DB.Create(&posts)
	
	if result.Error != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}
	// Return it
	ctx.JSON(http.StatusOK, gin.H{"messgae": "Post created successfully"})
}


func GetPosts(ctx *gin.Context){
	// Get posts
    var posts []models.Post
	initializers.DB.Find(&posts)

	// Respond with them
	ctx.JSON(http.StatusOK, gin.H{"Post": posts})

}