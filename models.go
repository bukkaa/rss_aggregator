package main

import (
	"time"

	"github.com/bukkaa/rss_aggregator/internal/database"
	"github.com/google/uuid"
)

type UserDto struct {
	ID        uuid.UUID `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	ApiKey    string    `json:"api_key"`
}

func newUserDto(entity database.User) UserDto {
	return UserDto{
		ID:        entity.ID,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
		Name:      entity.Name,
		ApiKey:    entity.ApiKey,
	}
}

type FeedDto struct {
	ID        uuid.UUID `json:"feed_id"`
	Name      string    `json:"name"`
	Url       string    `json:"url"`
	UserID    uuid.UUID `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func newFeedDto(entity database.Feed) FeedDto {
	return FeedDto{
		ID:        entity.ID,
		Name:      entity.Name,
		Url:       entity.Url,
		UserID:    entity.UserID,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}
}

func newListFeedDto(feedEntities []database.Feed) []FeedDto {
	var feedList []FeedDto
	for _, v := range feedEntities {
		feedList = append(feedList, newFeedDto(v))
	}

	return feedList
}

type FeedFollowsDto struct {
	ID        uuid.UUID `json:"id"`
	UserID    uuid.UUID `json:"user_id"`
	FeedID    uuid.UUID `json:"feed_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func newFeedFollowsDto(entity database.FeedFollow) FeedFollowsDto {
	return FeedFollowsDto{
		ID:        entity.ID,
		UserID:    entity.UserID,
		FeedID:    entity.FeedID,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}
}