package main

import (
	"dunky-star/go-simple/controllers"
	"dunky-star/go-simple/initializers"

	"github.com/gin-gonic/gin"
)

func init(){
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main(){

	r := gin.Default()

	r.POST("/api/v1/posts", controllers.PostCreate)
	r.GET("/api/v1/posts", controllers.GetAllPosts)
	r.GET("/api/v1/posts/:id", controllers.GetAPost)
	r.PUT("/api/v1/posts/:id", controllers.UpdatePost)
	r.DELETE("/api/v1/posts/:id", controllers.DeletePost)


	r.Run() // listen and serve on 0.0.0.0:9090 (for windows "localhost:9090")
  }
