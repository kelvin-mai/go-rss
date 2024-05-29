package model

type AuthInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type FeedInput struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}
