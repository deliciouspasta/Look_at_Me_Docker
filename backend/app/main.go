package main
 
import (
	// "user-api/user-api/routes"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
  	// "time"
)
 
func main() {
	// router := routes.NewRoutes()
	// router.Run(":8888")

	r := gin.Default()
    
    // CORS 対応
    config := cors.DefaultConfig()
    config.AllowOrigins = []string{"http://localhost:8080"}
    r.Use(cors.New(config))

    // CORS 対応の後にルーティングを書かないとうまく動かない
    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong",
        })
    })

    r.Run(":8888")
}