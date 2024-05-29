package model

import "time"

type AuthPayload struct {
	Username  string    `json:"username"`
	ExpiresAt time.Time `json:"expiresAt"`
}

type AuthInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
