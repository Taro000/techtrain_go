
package models

import (
	"time"
)

// UserCharacter represents a row from 'techtrain_game.user_character'.
type UserCharacter struct {
	UserID          string    `json:"user_id"`           // user_id
	CharacterID     string    `json:"character_id"`      // character_id
	UserCharacterID string    `json:"user_character_id"` // user_character_id
	CreatedAt       time.Time `json:"created_at"`        // created_at
	ModifiedAt      time.Time `json:"modified_at"`       // modified_at
}

