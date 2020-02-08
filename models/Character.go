
package models



import (
	"time"
)

// Character represents a row from 'techtrain_game.character'.
type Character struct {
	ID string `json:"id"` // id
	Name string `json:"name"` // name
	CreatedAt time.Time `json:"created_at"` // created_at
	ModifiedAt time.Time `json:"modified_at"` // modified_at

	
	

}


