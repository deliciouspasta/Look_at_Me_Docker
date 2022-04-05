package handlers
 
import (
	"net/http"
	"user-api/user-api/models"
	"time"
 
	"github.com/gin-gonic/gin"
	"github.com/rs/xid"

	"encoding/json"
	"io/ioutil"
)
 
// 一時的なtmpDB
// 将来的にはDatabaseに格納したい
var users []models.User
 
func init() {
	users = make([]models.User, 0)
	file, _ := ioutil.ReadFile("users.json")
	_ = json.Unmarshal([]byte(file), &users)
}
 
func NewUserHandler(c *gin.Context) {
	var user models.User
	// リクエストデータを取り出す
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H {
			"error": err.Error()})
		return
	}
 
	// ユニークなIDを生成
	user.ID = xid.New().String()
	// 現在時刻を追加
	user.PublishedAt = time.Now()
	users = append(users, user)
	c.JSON(http.StatusOK, user)
}