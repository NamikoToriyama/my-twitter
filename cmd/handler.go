package main

import (
	"log"
	"net/http"
	"time"

	"github.com/mercari/Week6/Week6/NamikoToriyama/app/domain"
)

// TweetHandler ... Interface of tweet_handler.go.
type TweetHandler interface {
	tweetListPage(http.ResponseWriter, *http.Request)
	tweetPage(http.ResponseWriter, *http.Request)
	tweetDetailPage(http.ResponseWriter, *http.Request)
	editPage(http.ResponseWriter, *http.Request)
	updateTweetPage(http.ResponseWriter, *http.Request)
	tweetDeletePage(http.ResponseWriter, *http.Request)
}

type tweetHandler struct {
	client domain.TweetDomain
}

// NewTweetHandler ... Return a structure that implements the interface.
func NewTweetHandler(s domain.TweetDomain) TweetHandler {
	return &tweetHandler{s}
}

// Tweet ... Structure for returning to html.
type Tweet struct {
	Username     string
	TweetContext string
	TweetID      string
	RegisterDate time.Time
}

func (tr *tweetHandler) tweetListPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tweets := tr.client.ListBlogs(w, r)
	contents := []Tweet{}
	for i := range tweets {
		data := Tweet{
			Username:     tweets[i].Username,
			TweetContext: tweets[i].Tweet,
			TweetID:      tweets[i].ID,
		}
		contents = append(contents, data)
	}
	tmpl.ExecuteTemplate(w, "tweet_list.html", contents)
}

func (tr *tweetHandler) tweetPage(w http.ResponseWriter, r *http.Request) {
	tr.client.CreateBlog(w, r)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := tmpl.ExecuteTemplate(w, "tweet.html", nil); err != nil {
		log.Fatal(err)
	}
}

func (tr *tweetHandler) tweetDetailPage(w http.ResponseWriter, r *http.Request) {
	tweet := tr.client.GetBlog(w, r)
	data := Tweet{
		Username:     tweet.Username,
		TweetContext: tweet.Tweet,
		TweetID:      tweet.ID,
		RegisterDate: tweet.RegisterDate,
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tmpl.ExecuteTemplate(w, "tweet_detail.html", data)
}

func (tr *tweetHandler) editPage(w http.ResponseWriter, r *http.Request) {
	tweet := tr.client.GetBlog(w, r)
	data := Tweet{
		Username:     tweet.Username,
		TweetContext: tweet.Tweet,
		TweetID:      tweet.ID,
		RegisterDate: tweet.RegisterDate,
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tmpl.ExecuteTemplate(w, "update_tweet.html", data)
}

func (tr *tweetHandler) updateTweetPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tweet := tr.client.EditBlog(w, r)
	data := Tweet{
		Username:     tweet.Username,
		TweetContext: tweet.Tweet,
		TweetID:      tweet.ID,
		RegisterDate: tweet.RegisterDate,
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tmpl.ExecuteTemplate(w, "tweet_detail.html", data)
}

func (tr *tweetHandler) tweetDeletePage(w http.ResponseWriter, r *http.Request) {
	tr.client.DeleteBlog(w, r)
}
