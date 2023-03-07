package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h handler) HomePage(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"Message": "This is home Route"})
}
