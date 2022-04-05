package routes
 
import (
	"user-api/user-api/handlers"
 
	"github.com/gin-gonic/gin"

	
)
 
func NewRoutes() *gin.Engine {
	router := gin.Default()
	router.POST("/users", handlers.NewUserHandler)
	
	router.GET("/users", handlers.ListUsersHandler)
	router.PUT("/users/:id", handlers.UpdateUserHandler)
	router.DELETE("/users/:id", handlers.DeleteUserHandler)
	router.GET("/users/search", handlers.SearchUsersHandler)
	return router
}