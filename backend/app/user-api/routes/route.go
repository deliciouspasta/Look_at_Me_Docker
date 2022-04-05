package routes
 
import (
	"user-api/user-api/handlers"
 
	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/cors"
	"time"

	
)
 
func NewRoutes() *gin.Engine {
	router := gin.Default()

	// config := cors.DefaultConfig()
    // config.AllowOrigins = []string{"http://localhost:8080"}
    // router.Use(cors.New(config))



	router.Use(cors.New(cors.Config{
		// アクセスを許可したいアクセス元
		AllowOrigins: []string{
			"http://localhost:8080",
		},
		// アクセスを許可したいHTTPメソッド
		AllowMethods: []string{
			"POST",
			"GET",
			"PUT",
			"DELETE",
			"OPTIONS",
		},
		// 許可したいHTTPリクエストヘッダ
		AllowHeaders: []string{
			"Access-Control-Allow-Credentials",
			"Access-Control-Allow-Headers",
			// "Access-Control-Allow-Origin",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"Authorization",
		},
		// cookieなどの情報を必要とするかどうか
		AllowCredentials: true,
		// preflightリクエストの結果をキャッシュする時間
		MaxAge: 24 * time.Hour,
	  }))


	router.POST("/users", handlers.NewUserHandler)
	
	router.GET("/users", handlers.ListUsersHandler)
	router.PUT("/users/:id", handlers.UpdateUserHandler)
	router.DELETE("/users/:id", handlers.DeleteUserHandler)
	router.GET("/users/search", handlers.SearchUsersHandler)
	return router
}