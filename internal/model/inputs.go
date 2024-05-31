package model

type AuthInput struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type FeedInput struct {
	Name string `json:"name" validate:"required"`
	Url  string `json:"url" validate:"required,url"`
}

type ValidationError struct {
	Field string      `json:"field"`
	Tag   string      `json:"tag"`
	Value interface{} `json:"value"`
}
