
package models

import (
	"time"
)

// Character represents a row from 'techtrain_game.characters'.
type Character struct {
	ID          string    `json:"id"`          // id
	Name        string    `json:"name"`        // name
	Probability int      `json:"probability"` // probability
	CreatedAt   time.Time `json:"created_at"`  // created_at
	ModifiedAt  time.Time `json:"modified_at"` // modified_at
	UserCharacters []UserCharacter
}

