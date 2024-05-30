package service

import (
	"github.com/jmoiron/sqlx"
	"kelvinmai.io/rss/internal/model"
)

type FeedService struct {
	db *sqlx.DB
}

func NewFeedService(db *sqlx.DB) *FeedService {
	return &FeedService{
		db: db,
	}
}

func (s *FeedService) GetAll() ([]model.Feed, error) {
	feeds := []model.Feed{}
	err := s.db.Select(&feeds, "select * from feeds")
	return feeds, err
}

func (s *FeedService) GetById(id string) (model.Feed, error) {
	feed := model.Feed{}
	err := s.db.Get(&feed, "select * from feeds where id = $1", id)
	return feed, err
}

func (s *FeedService) Create(name, url string) (*model.Feed, error) {
	rows, err := s.db.Queryx(
		`insert into feeds (name, url)
		 values ($1, $2)
		 returning *`,
		name,
		url,
	)
	if err != nil {
		return nil, err
	}
	feed := model.Feed{}
	rows.Next()
	err = rows.StructScan(&feed)
	return &feed, err
}

func (s *FeedService) Update(id, name, url string) (*model.Feed, error) {
	rows, err := s.db.Queryx(
		`update feeds
		 set name = $2,
		     url = $3
		 where id = $1
		 returning *`,
		id,
		name,
		url,
	)
	if err != nil {
		return nil, err
	}
	feed := model.Feed{}
	rows.Next()
	err = rows.StructScan(&feed)
	return &feed, err
}

func (s *FeedService) Delete(id string) (*model.Feed, error) {
	rows, err := s.db.Queryx(
		`delete from feeds
		 where id = $1
		 returning *`,
		id,
	)
	if err != nil {
		return nil, err
	}
	feed := model.Feed{}
	rows.Next()
	err = rows.StructScan(&feed)
	return &feed, err
}

func (s *FeedService) Follow(id, userId string) (*model.UserFeed, error) {
	rows, err := s.db.Queryx(
		`insert into user_feed
		 (feed_id, user_id)
		 values
		 ($1, $2)
		 returning *`,
		id,
		userId,
	)
	if err != nil {
		return nil, err
	}
	userFeed := model.UserFeed{}
	rows.Next()
	err = rows.StructScan(&userFeed)
	return &userFeed, err
}

func (s *FeedService) Unfollow(id, userId string) (*model.UserFeed, error) {
	rows, err := s.db.Queryx(
		`delete from user_feed
		 where feed_id = $1
		   and user_id = $2
		 returning *`,
		id,
		userId,
	)
	if err != nil {
		return nil, err
	}
	userFeed := model.UserFeed{}
	rows.Next()
	err = rows.StructScan(&userFeed)
	return &userFeed, err
}

func (s *FeedService) GetUserFeeds(userId string) (interface{}, error) {
	feeds := []model.Feed{}
	err := s.db.Select(
		&feeds,
		`select feeds.*
		 from user_feed
		 join feeds on feeds.id = user_feed.feed_id
		 where user_id = $1`,
		userId,
	)
	if err != nil {
		return nil, err
	}
	return feeds, err
}
