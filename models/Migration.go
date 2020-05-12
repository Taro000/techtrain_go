
package models

import (
	"time"
)

// Migration represents a row from 'techtrain_game.migrations'.
type Migration struct {
	ID        string         `json:"id"`         // id
	AppliedAt *time.Time `json:"applied_at"` // applied_at
}

