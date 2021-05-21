package domain

import (
	"net/http"

	"github.com/mercari/Week6/Week6/NamikoToriyama/app/service"
	"github.com/mercari/Week6/Week6/NamikoToriyama/model"
	"github.com/rs/xid"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

// TweetDomain ... Interface of tweet_domain.go.
type TweetDomain interface {
	CreateBlog(http.ResponseWriter, *http.Request)
	DeleteBlog(http.ResponseWriter, *http.Request)
	EditBlog(http.ResponseWriter, *http.Request) *model.Tweet
	GetBlog(http.ResponseWriter, *http.Request) *model.Tweet
	ListBlogs(http.ResponseWriter, *http.Request) []*model.Tweet
}

type tweetDomain struct {
	client service.TweetService
}

// NewTweetDomain ... Return a structure that implements the interface.
func NewTweetDomain(s service.TweetService) TweetDomain {
	return &tweetDomain{s}
}

// ListBlogs ... Return a list of tweets.
func (th *tweetDomain) ListBlogs(res http.ResponseWriter, req *http.Request) []*model.Tweet {
	ctx := appengine.NewContext(req)
	switch req.Method {
	case http.MethodGet:
		// Get list of blogs from file or database
		tweetList, err := th.client.ListTweet(ctx)
		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
			log.Errorf(ctx, "Fail to post tweet!: %v", err)
			return nil
		}
		res.WriteHeader(http.StatusOK)
		return tweetList
	default:
		res.WriteHeader(http.StatusMethodNotAllowed)
		log.Errorf(ctx, "405 - Method not allowed!")
	}
	return nil
}

// CreateBlog ... Modify form to tweet structure and create a Tweet.
func (th *tweetDomain) CreateBlog(res http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)
	switch req.Method {
	case http.MethodPost:
		err := req.ParseForm()

		if err != nil {
			res.WriteHeader(http.StatusBadRequest)
			log.Errorf(ctx, "Could not parse form data!: %v", err)
			return
		}

		// this is just dummy data remove this
		tweet := new(model.Tweet)
		guid := xid.New()
		tweet.ID = guid.String()
		tweet.Username = req.Form.Get("username")
		tweet.Tweet = req.Form.Get("tweet")

		// store data of the request in a file or database
		if th.client.PostTweet(ctx, tweet) != nil {
			res.WriteHeader(http.StatusInternalServerError)
			log.Errorf(ctx, "Fail to post tweet!: %v", err)
			return
		}

		// Return blog data to client in json format
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusCreated)

	default:
		res.WriteHeader(http.StatusMethodNotAllowed)
		log.Errorf(ctx, "405 - Method not allowed!")
	}

}

// GetBlog ... Return a tweets.
func (th *tweetDomain) GetBlog(res http.ResponseWriter, req *http.Request) *model.Tweet {
	ctx := appengine.NewContext(req)
	switch req.Method {
	case http.MethodGet:
		// Get tweetID from the query parameters (?tweetID=...)
		queryParameters := req.URL.Query()
		tweetID := queryParameters.Get("tweetID")

		// Get data of blog with id from parameters from file or database
		tweet, err := th.client.GetTweet(ctx, tweetID)
		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
			log.Errorf(ctx, "Fail to get tweet!: %v", err)
			return nil
		}

		// Return blog data to client in json format
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusOK)
		return tweet

	default:
		res.WriteHeader(http.StatusMethodNotAllowed)
		log.Warningf(ctx, "405 - Method not allowed!")
	}
	return nil
}

// DeleteBlog ... Delete the tweet with tweetID
func (th *tweetDomain) DeleteBlog(res http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)

	switch req.Method {
	case http.MethodDelete:
		// Get tweetID from the query parameters (?tweetID=...)
		queryParameters := req.URL.Query()
		tweetID := queryParameters.Get("tweetID")

		// delete data with blogId from request in a file or database
		if err := th.client.DeleteTweet(ctx, tweetID); err != nil {
			res.WriteHeader(http.StatusInternalServerError)
			log.Errorf(ctx, "Fail to get tweet!: %v", err)
			return
		}

		// Return empty response to client in json format
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusOK)
	default:
		res.WriteHeader(http.StatusMethodNotAllowed)
		log.Warningf(ctx, "405 - Method not allowed!")
	}
}

// EditBlog ...  Modify form to tweet structure and update a Tweet.
func (th *tweetDomain) EditBlog(res http.ResponseWriter, req *http.Request) *model.Tweet {
	ctx := appengine.NewContext(req)
	switch req.Method {
	case http.MethodPost:
		err := req.ParseForm()
		if err != nil {
			res.WriteHeader(http.StatusBadRequest)
			log.Errorf(ctx, "Could not parse form data!: %v", err)

			return nil
		}
		tweet := new(model.Tweet)
		tweet.ID = req.Form.Get("tweetID")
		tweet.Username = req.Form.Get("username")
		tweet.Tweet = req.Form.Get("tweet")
		updateTweet, err := th.client.UpdateTweet(ctx, tweet)
		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
			log.Errorf(ctx, "Fail to update tweet!: %v", err)
			return nil
		}

		// Return updated blog data to client in json format
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusCreated)
		return updateTweet
	default:
		res.WriteHeader(http.StatusMethodNotAllowed)
		log.Warningf(ctx, "405 - Method not allowed!")

	}
	return nil
}
