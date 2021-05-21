package repository

import (
	"context"

	"github.com/mercari/Week6/Week6/NamikoToriyama/model"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
)

// TweetRepository ...Interface of tweet_repository.go.
type TweetRepository interface {
	List(context.Context) ([]*model.Tweet, error)
	Put(context.Context, *model.Tweet) (int, error)
	Get(context.Context, string) (*model.Tweet, error)
	Delete(context.Context, string) error
}

type tweetRepository struct{}

// NewTweetRepository ... Return a structure that implements the interface.
func NewTweetRepository() TweetRepository {
	return &tweetRepository{}
}

// List ...  Get a list from Google cloud datastore.
func (t *tweetRepository) List(ctx context.Context) ([]*model.Tweet, error) {
	tweets := []*model.Tweet{}
	keys, err := datastore.NewQuery("Tweet").Order("-registerDate").GetAll(ctx, &tweets)
	if err != nil {
		log.Errorf(ctx, "Fail to get category list: %v", err)
		return nil, err
	}
	log.Infof(ctx, "Success to get category list")
	for i := range tweets {
		log.Infof(ctx, "%#v", keys[i].String())
		log.Infof(ctx, "%#v", tweets[i])
	}

	return tweets, nil
}

// Put ... Put a tweet in Google cloud datastore.
func (t *tweetRepository) Put(ctx context.Context, tweet *model.Tweet) (int, error) {
	key := t.createKey(ctx, tweet.ID)

	key, err := datastore.Put(ctx, key, tweet)
	if err != nil {
		log.Errorf(ctx, "Fail to put category: %v", err)
		return 0, err
	}

	log.Infof(ctx, "Success to put tweet")
	log.Infof(ctx, "id: %v", key.IntID())

	return int(key.IntID()), nil
}

// Get ... Get a tweet from Google cloud datastore.
func (t *tweetRepository) Get(ctx context.Context, id string) (*model.Tweet, error) {
	key := t.createKey(ctx, id)

	tweet := new(model.Tweet)
	if err := datastore.Get(ctx, key, tweet); err != nil {
		log.Errorf(ctx, "Fail to get Tweet: %v", err)
		return nil, err
	}

	tweet.ID = id
	log.Debugf(ctx, "Success to get Tweet")
	log.Debugf(ctx, "%#v", tweet)

	return tweet, nil
}

// Delete ... Delete tweets in Google cloud datastore.
func (t *tweetRepository) Delete(ctx context.Context, id string) error {
	key := t.createKey(ctx, id)

	if err := datastore.Delete(ctx, key); err != nil {
		log.Errorf(ctx, "Fail to delete Tweet: %v", err)
		return err
	}
	log.Infof(ctx, "Success to delete Tweet")

	return nil
}

func (t *tweetRepository) createKey(ctx context.Context, id string) *datastore.Key {
	return datastore.NewKey(ctx, "Tweet", id, 0, nil)
}
