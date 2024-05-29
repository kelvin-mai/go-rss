package controller

import (
	"github.com/gofiber/fiber/v2"
	"kelvinmai.io/rss/internal/model"
	"kelvinmai.io/rss/internal/router/response"
	"kelvinmai.io/rss/internal/service"
)

type FeedController struct {
	s *service.FeedService
}

func NewFeedController(s *service.FeedService) *FeedController {
	return &FeedController{
		s: s,
	}
}

func (c *FeedController) GetAllFeeds(ctx *fiber.Ctx) error {
	feeds, err := c.s.GetAll()
	if err != nil {
		return err
	}
	return response.Ok(ctx, fiber.Map{
		"feeds": feeds,
	})
}

func (c *FeedController) GetFeedById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	feed, err := c.s.GetById(id)
	if err != nil {
		return err
	}
	return response.Ok(ctx, fiber.Map{
		"feed": feed,
	})
}

func (c *FeedController) CreateFeed(ctx *fiber.Ctx) error {
	input := model.FeedInput{}
	if err := ctx.BodyParser(&input); err != nil {
		return err
	}
	feed, err := c.s.Create(input.Name, input.Url)
	if err != nil {
		return err
	}

	return response.Created(ctx, fiber.Map{
		"feed": feed,
	})
}

func (c *FeedController) UpdateFeed(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	input := model.FeedInput{}
	if err := ctx.BodyParser(&input); err != nil {
		return err
	}
	feed, err := c.s.Update(id, input.Name, input.Url)
	if err != nil {
		return err
	}
	return response.Ok(ctx, fiber.Map{
		"feed": feed,
	})
}

func (c *FeedController) DeleteById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	feed, err := c.s.Delete(id)
	if err != nil {
		return err
	}
	return response.Ok(ctx, fiber.Map{
		"feed": feed,
	})
}
