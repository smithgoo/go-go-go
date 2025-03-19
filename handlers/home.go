package handlers

import (
	"github.com/gin-gonic/gin"
	//"go-go-go/models"
	"net/http"
)

func Home(c *gin.Context) {
	//userID := c.MustGet("userID").(uint)
	//var user models.User
	//if err := models.DB.First(&user, userID).Error; err != nil {
	//	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	//	return
	//}
	c.JSON(http.StatusOK, gin.H{"message": "success!"})
}
