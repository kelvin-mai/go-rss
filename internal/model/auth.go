package model

import "time"

type AuthPayload struct {
	Username  string    `json:"username"`
	UserId    string    `json:"userId"`
	IsAdmin   bool      `json:"isAdmin"`
	ExpiresAt time.Time `json:"expiresAt"`
}
