package handlers

import (
	"SpotifyMusicoo/Models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h handler) GetUserById(c *gin.Context) {
	id := c.Param("id")
	var user Models.User
	if result := h.DB.Find(&user, id); result.Error != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"Message": "UserNot Found"})
	}
	c.IndentedJSON(http.StatusOK, &user)
}
