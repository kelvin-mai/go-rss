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
	err := s.db.Select(&user, "select * from users where id = $1", id)
	return user, err
}
