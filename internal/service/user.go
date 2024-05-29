package service

import (
	"github.com/jmoiron/sqlx"
	"kelvinmai.io/rss/internal/model"
)

type UserService struct {
	db *sqlx.DB
}

func NewUserService(db *sqlx.DB) *UserService {
	return &UserService{
		db: db,
	}
}

func (s *UserService) GetAll() ([]model.User, error) {
	users := []model.User{}
	err := s.db.Select(&users, "select * from users")
	return users, err
}

func (s *UserService) GetById(id string) (model.User, error) {
	user := model.User{}
	err := s.db.Get(&user, "select * from users where id = $1", id)
	return user, err
}

func (s *UserService) GetByUsername(username string) (model.User, error) {
	user := model.User{}
	err := s.db.Get(&user, "select * from users where username = $1", username)
	return user, err
}

func (s *UserService) Create(username, password string) error {
	_, err := s.db.Exec(
		`insert into users (username, password)
		 values ($1, $2)
		`,
		username,
		password,
	)
	return err
}
