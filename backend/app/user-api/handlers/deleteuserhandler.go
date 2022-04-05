package handlers
 
import (
	"net/http"
 
	"github.com/gin-gonic/gin"
)
 
func DeleteUserHandler(c *gin.Context) {
	// リクエストからIDを取得
	id := c.Param("id")
 
	userIndex := -1
	for i := 0; i < len(users); i++ {
		if users[i].ID == id {
			userIndex = i
		}
	}
 
	// レシピが見つからない場合
	if userIndex == -1 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "user not found"})
		return
	}
 
	// idが一致したレシピ"以外"をusers変数に詰め直す
	users = append(users[:userIndex], users[userIndex+1:]...)
 
	c.JSON(http.StatusOK, gin.H{
		"message": "User has been deleted"})
}