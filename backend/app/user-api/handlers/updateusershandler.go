package handlers
 
import (
	"net/http"
	"user-api/user-api/models"
 
	"github.com/gin-gonic/gin"
)
 
func UpdateUserHandler(c *gin.Context) {
	// urlに指定するusers/:idからidを取得できる
	id := c.Param("id")
	var user models.User
 
	if err := c.ShouldBindJSON(&user); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
 
	targetUserIndex := -1
	for i := 0; i < len(users); i++ {
		if users[i].ID == id {
			targetUserIndex = i
		}
	}
 
	if targetUserIndex == -1 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "User not found"})
		return
	}
 
	users[targetUserIndex] = user
	c.JSON(http.StatusOK, user)
}