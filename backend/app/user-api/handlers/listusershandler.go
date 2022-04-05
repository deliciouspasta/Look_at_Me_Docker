package handlers
 
import (
	"net/http"
 
	"github.com/gin-gonic/gin"
)
 
func ListUsersHandler(c *gin.Context) {
	c.JSON(http.StatusOK, users)
}