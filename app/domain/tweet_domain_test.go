package domain

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/mercari/Week6/Week6/NamikoToriyama/mock_app"
	"github.com/mercari/Week6/Week6/NamikoToriyama/model"
)

func TestListBlog(t *testing.T) {
	now := time.Now()
	want := []*model.Tweet{
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

	mock := mock_app.NewMockTweetService(ctrl)
	mock.EXPECT().ListTweet(gomock.Any()).Return(want, nil)
	rr := httptest.NewRecorder()
	target := &tweetDomain{
		client: mock,
	}

	// Execute test.
	got := target.ListBlogs(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf(
			"unexpected status: got (%v) want (%v)",
			status,
			http.StatusOK,
		)
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("want = %v, got = %v", want, got)
	}

}

func TestCreateBlog(t *testing.T) {
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

	mock := mock_app.NewMockTweetService(ctrl)
	mock.EXPECT().PostTweet(gomock.Any(), gomock.Any()).Return(nil)
	rr := httptest.NewRecorder()
	target := &tweetDomain{
		client: mock,
	}

	// Execute test.
	handler := http.HandlerFunc(target.CreateBlog)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf(
			"unexpected status: got (%v) want (%v)",
			status,
			http.StatusCreated,
		)
	}
}

func TestGetBlog(t *testing.T) {
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

	mock := mock_app.NewMockTweetService(ctrl)
	mock.EXPECT().GetTweet(gomock.Any(), gomock.Any()).Return(want, nil)
	rr := httptest.NewRecorder()

	target := &tweetDomain{
		client: mock,
	}

	// Execute test.
	got := target.GetBlog(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf(
			"unexpected status: got (%v) want (%v)",
			status,
			http.StatusOK,
		)
	}
	if got != want {
		t.Fatalf("want = %v, got = %v", want, got)
	}
}

func TestEditBlog(t *testing.T) {
	now := time.Now()
	want := &model.Tweet{
		ID:           "12345",
		Username:     "foo",
		Tweet:        "Hello world!!",
		RegisterDate: now,
	}
	values := url.Values{} // url.Valuesオブジェクト生成
	values.Add("ID", "12345")
	values.Add("Username", "foo")
	values.Add("Tweet", "Hello world!!")

	req, err := http.NewRequest("POST", "/edit", strings.NewReader(values.Encode()))
	if err != nil {
		t.Fatal(err)
	}

	// Set return value of mock.
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock_app.NewMockTweetService(ctrl)
	mock.EXPECT().UpdateTweet(gomock.Any(), gomock.Any()).Return(want, nil)
	rr := httptest.NewRecorder()

	target := &tweetDomain{
		client: mock,
	}

	// Execute test.
	got := target.EditBlog(rr, req)
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf(
			"unexpected status: got (%v) want (%v)",
			status,
			http.StatusCreated,
		)
	}
	if got != want {
		t.Fatalf("want = %v, got = %v", want, got)
	}
}
func TestDeleteBlog(t *testing.T) {

	req, err := http.NewRequest("DELETE", "/delete?tweetID=12345", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Set return value of mock.
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock_app.NewMockTweetService(ctrl)
	mock.EXPECT().DeleteTweet(gomock.Any(), gomock.Any()).Return(nil)

	rr := httptest.NewRecorder()

	target := &tweetDomain{
		client: mock,
	}

	// Execute test.
	handler := http.HandlerFunc(target.DeleteBlog)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf(
			"unexpected status: got (%v) want (%v)",
			status,
			http.StatusOK,
		)
	}

}
