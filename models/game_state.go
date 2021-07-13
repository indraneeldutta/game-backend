package models

import "time"

type GameState struct {
	UserID      string    `bson:"user_id" json:"id,omitempty"`
	GamesPlayed int64     `json:"gamesPlayed"`
	Score       int64     `json:"score"`
	CreatedAt   time.Time `bson:"created_at" json:"created_at,omitempty"`
}
