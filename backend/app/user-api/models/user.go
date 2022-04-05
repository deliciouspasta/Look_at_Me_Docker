package models
 
import "time"
 
type User struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	Email        string  `json:"email"`
	Contents  []string  `json:"contents"`
	Tags []string  `json:"tags"`
	PublishedAt  time.Time `json:"publishedAt"`
}