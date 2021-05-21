package service

import (
	"context"
	"time"

	"github.com/mercari/Week6/Week6/NamikoToriyama/app/repository"
	"github.com/mercari/Week6/Week6/NamikoToriyama/model"
)

// TweetService ... Interface of tweet_service.go.
type TweetService interface {
	ListTweet(context.Context) ([]*model.Tweet, error)
	PostTweet(context.Context, *model.Tweet) error
	GetTweet(context.Context, string) (*model.Tweet, error)
	UpdateTweet(context.Context, *model.Tweet) (*model.Tweet, error)
	DeleteTweet(context.Context, string) error
}

type tweetService struct {
	client repository.TweetRepository
}

// NewTweetService ... Return a structure that implements the interface.
func NewTweetService(s repository.TweetRepository) TweetService {
	return &tweetService{s}
}

//ListTweet ... Return a list of tweets.
func (t *tweetService) ListTweet(ctx context.Context) ([]*model.Tweet, error) {
	return t.client.List(ctx)
}

// PostTweet ... Create a Tweet.
func (t *tweetService) PostTweet(ctx context.Context, tweet *model.Tweet) error {
	now := time.Now()
	tweet.RegisterDate = now

	if _, err := t.client.Put(ctx, tweet); err != nil {
		return err
	}

	return nil
}

// GetTweet ... Get a Tweet.
func (t *tweetService) GetTweet(ctx context.Context, id string) (*model.Tweet, error) {
	tweet, err := t.client.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	tweet.ID = id

	return tweet, nil
}

// UpdateTweet ... Update a Tweet.
func (t *tweetService) UpdateTweet(ctx context.Context, tweet *model.Tweet) (*model.Tweet, error) {
	exTweet, err := t.client.Get(ctx, tweet.ID)

	if err != nil {
		return nil, err
	}

	exTweet.Username = tweet.Username
	exTweet.Tweet = tweet.Tweet
	if _, err := t.client.Put(ctx, exTweet); err != nil {
		return nil, err
	}

	return tweet, nil
}

// DeleteTweet ... Delete a Tweet.
func (t *tweetService) DeleteTweet(ctx context.Context, id string) error {
	return t.client.Delete(ctx, id)
}
