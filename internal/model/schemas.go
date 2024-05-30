package model

import (
	"time"
)

type User struct {
	Id        string     `json:"id"`
	CreatedAt time.Time  `db:"created_at" json:"createdAt"`
	UpdatedAt *time.Time `db:"updated_at" json:"updatedAt"`
	Username  string     `json:"username"`
	Password  string     `json:"-"`
	IsAdmin   bool       `db:"is_admin" json:"isAdmin"`
}

type Feed struct {
	Id        string     `json:"id"`
	CreatedAt time.Time  `db:"created_at" json:"createdAt"`
	UpdatedAt *time.Time `db:"updated_at" json:"updatedAt"`
	Name      string     `json:"name"`
	Url       string     `json:"url"`
}

type UserFeed struct {
	Id        string     `json:"id"`
	CreatedAt time.Time  `db:"created_at" json:"createdAt"`
	UpdatedAt *time.Time `db:"updated_at" json:"updatedAt"`
	UserId    string     `db:"user_id" json:"userId"`
	FeedId    string     `db:"feed_id" json:"feedId"`
}
