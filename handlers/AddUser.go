package handlers

import (
	"SpotifyMusicoo/Models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h handler) AddUser(c *gin.Context) {
	var NewUser = Models.User{}
	err := c.BindJSON(&NewUser)
	if err != nil {
		c.IndentedJSON(http.StatusUnprocessableEntity, gin.H{"Message": "Format of the Request is wrong"})
		return
	}
	if result := h.DB.Create(&NewUser); result.Error != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"Message": "Unknown Error"})
		return
	}
	c.IndentedJSON(http.StatusCreated, NewUser)

}
