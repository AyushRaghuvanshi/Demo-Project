package main

import (
	"SpotifyMusicoo/Database"
	"SpotifyMusicoo/handlers"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	DB := Database.Init()
	h := handlers.New(DB)
	gin := gin.Default()
	gin.GET("/", h.HomePage)
	gin.POST("/register", h.AddUser)
	gin.DELETE("/user/:id", h.DeleteUserById)
	gin.GET("/user/:id", h.GetUserById)
	gin.Run("localhost:8080")
}
func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
