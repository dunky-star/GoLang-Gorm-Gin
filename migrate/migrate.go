package main

import (
	"dunky-star/go-simple/initializers"
	"dunky-star/go-simple/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main(){
	initializers.DB.AutoMigrate(&models.Post{})
}
