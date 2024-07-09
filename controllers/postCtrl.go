package controllers

import (
	"dunky-star/go-simple/initializers"
	"dunky-star/go-simple/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type postRequestData struct {
	Title string `json:"title" binding:"required"`
	Body string `json:"body" binding:"required"`
}

// type postResponse struct {
// 	Title string `json:"title"`
// 	Body string  `json:"body"`
// }


func PostCreate(ctx *gin.Context) {
	// Get data off request body
    var req postRequestData

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
	ctx.JSON(http.StatusOK, gin.H{"message": "Post created successfully"})
}


func GetAllPosts(ctx *gin.Context){
	// Get posts
    var posts []models.Post
	initializers.DB.Find(&posts)

	// Respond with them
	ctx.JSON(http.StatusOK, gin.H{"Post": posts})

}

// Get post by ID request handler
type getPostID struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func GetAPost(ctx *gin.Context){
   // Get id off request URL
   var req getPostID

   if err := ctx.ShouldBindUri(&req); err != nil {
	ctx.JSON(http.StatusBadRequest, err)
	return
   }

	// Get posts
	var post models.Post
	initializers.DB.First(&post, req.ID)

	// Respond with them
	ctx.JSON(http.StatusOK, gin.H{"Post": post})

}

func UpdatePost (ctx *gin.Context){
	// Get id off request Url
    var req getPostID

	if err := ctx.ShouldBindUri(&req); err != nil{
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	// Get the data off request body
	var arg postRequestData
	if err := ctx.ShouldBindJSON(&arg); err != nil{
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
    
	// Find the post we are updating
    var post models.Post
	initializers.DB.Find(&post, req.ID)

	// Update it
    initializers.DB.Model(&post).Updates(models.Post{
		Title: arg.Title,
		Body: arg.Body,
	})
	// Response with status
	ctx.JSON(http.StatusOK, gin.H{"message": "Post updated successfully"})

}