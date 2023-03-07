package handlers

import (
	"SpotifyMusicoo/Models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h handler) DeleteUserById(c *gin.Context) {
	id := c.Param("id")
	var user Models.User
	if result := h.DB.Find(&user, id); result.Error != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"Message": "Unable to delete"})
		return
	}
	h.DB.Delete(&user)
	c.IndentedJSON(http.StatusOK, gin.H{"Message": "Deleted"})
}
