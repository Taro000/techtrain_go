
package models



import (
	"time"
)

// User represents a row from 'techtrain_game.user'.
type User struct {
	ID string `json:"id"` // id
	Name string `json:"name"` // name
	Token string `json:"token"` // token
	CreatedAt time.Time `json:"created_at"` // created_at
	ModifiedAt time.Time `json:"modified_at"` // modified_at

	
	

}


