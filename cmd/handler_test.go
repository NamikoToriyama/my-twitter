package main

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/mercari/Week6/Week6/NamikoToriyama/mock_app"
	"github.com/mercari/Week6/Week6/NamikoToriyama/model"
)

func TestListBlog(t *testing.T) {
	now := time.Now()
	tweetList := []*model.Tweet{
		{
			ID:           "001",
			Username:     "foo",
			Tweet:        "Good morning",
			RegisterDate: now,
		},
		{
			ID:           "002",
			Username:     "hoge",
			Tweet:        "Good afternoon",
			RegisterDate: now,
		},
		{
			ID:           "003",
			Username:     "bar",
			Tweet:        "Good evening",
			RegisterDate: now,
		},
	}

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	// Set return value of mock.
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock_app.NewMockTweetDomain(ctrl)
	mock.EXPECT().ListBlogs(gomock.Any(), gomock.Any()).Return(tweetList)
	rr := httptest.NewRecorder()
	target := &tweetHandler{
		client: mock,
	}

	// Execute test.
	handler := http.HandlerFunc(target.tweetListPage)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf(
			"unexpected status: got (%v) want (%v)",
			status,
			http.StatusOK,
		)
	}

}

func TestTweetPage(t *testing.T) {
	// Set form parameter
	values := url.Values{}
	values.Add("ID", "12345")
	values.Add("Username", "foo")
	values.Add("Tweet", "hello")
	req, err := http.NewRequest("POST", "/tweet", strings.NewReader(values.Encode()))
	if err != nil {
		t.Fatal(err)
	}

	// Set return value of mock.
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock_app.NewMockTweetDomain(ctrl)
	mock.EXPECT().CreateBlog(gomock.Any(), gomock.Any()).Return()
	rr := httptest.NewRecorder()
	target := &tweetHandler{
		client: mock,
	}

	// Execute test.
	handler := http.HandlerFunc(target.tweetPage)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf(
			"unexpected status: got (%v) want (%v)",
			status,
			http.StatusOK,
		)
	}
}

func TestTweetDetailPage(t *testing.T) {
	now := time.Now()
	want := &model.Tweet{
		ID:           "12345",
		Username:     "foo",
		Tweet:        "hello world",
		RegisterDate: now,
	}

	req, err := http.NewRequest("GET", "/get?tweetID=12345", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Set return value of mock.
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock_app.NewMockTweetDomain(ctrl)
	mock.EXPECT().GetBlog(gomock.Any(), gomock.Any()).Return(want)
	rr := httptest.NewRecorder()

	target := &tweetHandler{
		client: mock,
	}

	// Execute test.
	handler := http.HandlerFunc(target.tweetDetailPage)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf(
			"unexpected status: got (%v) want (%v)",
			status,
			http.StatusOK,
		)
	}
}

func TestEditPage(t *testing.T) {
	now := time.Now()
	want := &model.Tweet{
		ID:           "12345",
		Username:     "foo",
		Tweet:        "hello world",
		RegisterDate: now,
	}

	req, err := http.NewRequest("GET", "/edit", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Set return value of mock.
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock_app.NewMockTweetDomain(ctrl)
	mock.EXPECT().GetBlog(gomock.Any(), gomock.Any()).Return(want)
	rr := httptest.NewRecorder()

	target := &tweetHandler{
		client: mock,
	}

	// Execute test.
	handler := http.HandlerFunc(target.editPage)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf(
			"unexpected status: got (%v) want (%v)",
			status,
			http.StatusOK,
		)
	}
}

func TestUpdateTweetPage(t *testing.T) {
	now := time.Now()
	want := &model.Tweet{
		ID:           "12345",
		Username:     "foo",
		Tweet:        "hello world",
		RegisterDate: now,
	}

	req, err := http.NewRequest("GET", "/update", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Set return value of mock.
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock_app.NewMockTweetDomain(ctrl)
	mock.EXPECT().EditBlog(gomock.Any(), gomock.Any()).Return(want)
	rr := httptest.NewRecorder()

	target := &tweetHandler{
		client: mock,
	}

	// Execute test.
	handler := http.HandlerFunc(target.updateTweetPage)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf(
			"unexpected status: got (%v) want (%v)",
			status,
			http.StatusOK,
		)
	}
}

func TestTweetDeletePage(t *testing.T) {

	req, err := http.NewRequest("DELETE", "/delete?tweetID=12345", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Set return value of mock.
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock_app.NewMockTweetDomain(ctrl)
	mock.EXPECT().DeleteBlog(gomock.Any(), gomock.Any()).Return()

	rr := httptest.NewRecorder()
	target := &tweetHandler{
		client: mock,
	}

	// Execute test.
	handler := http.HandlerFunc(target.tweetDeletePage)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf(
			"unexpected status: got (%v) want (%v)",
			status,
			http.StatusOK,
		)
	}

}
