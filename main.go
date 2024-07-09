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
	r.GET("/api/v1/posts", controllers.GetPosts)

	r.Run() // listen and serve on 0.0.0.0:9090 (for windows "localhost:9090")
  }
