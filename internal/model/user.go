package model

import (
	"time"
)

type User struct {
	Id        string     `db:"id" json:"id"`
	CreatedAt time.Time  `db:"created_at" json:"createdAt"`
	UpdatedAt *time.Time `db:"updated_at" json:"updatedAt"`
	Username  string     `db:"username" json:"username"`
	Password  string     `db:"password" json:"password"`
}
