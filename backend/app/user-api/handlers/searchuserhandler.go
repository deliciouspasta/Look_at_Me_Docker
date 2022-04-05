package handlers
 
import (
	"net/http"
	"user-api/user-api/models"
	"strings"
 
	"github.com/gin-gonic/gin"
)
 
func SearchUsersHandler(c *gin.Context) {

	tag := c.Query("tags")
 
	listOfUsers := make([]models.User, 0)
 
	for i := 0; i < len(users); i++ {
		found := false
		for _, t := range users[i].Tags {
			if strings.EqualFold(t, tag) {
				found = true
			}
		}
		if found {
			listOfUsers = append(listOfUsers, users[i])
		}
	}
	c.JSON(http.StatusOK, listOfUsers)
}