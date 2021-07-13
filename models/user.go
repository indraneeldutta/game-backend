package models

import "time"

type CreateUser struct {
	ID        string    `bson:"_id" json:"id,omitempty"`
	Name      string    `json:"name"`
	CreatedAt time.Time `bson:"created_at" json:"created_at,omitempty"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at,omitempty"`
}

type CreateUserResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
