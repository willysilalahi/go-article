package response

import (
	"time"
)

type PostResponse struct {
	ID          int                      `json:"id"`
	Title       string                   `json:"title"`
	Description string                   `json:"description" gorm:"type:text"`
	Status      int                      `json:"status"`
	AdminID     int                      `json:"admin_id"`
	Admin       map[string]interface{}   `json:"admin"`
	Comments    []map[string]interface{} ` json:"comments"`
	CreatedAt   time.Time                `json:"created_at"`
}
