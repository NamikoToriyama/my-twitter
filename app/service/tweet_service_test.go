package service

import (
	"context"
	"reflect"
	"testing"
	"time"

	gomock "github.com/golang/mock/gomock"
	"github.com/mercari/Week6/Week6/NamikoToriyama/mock_app"
	"github.com/mercari/Week6/Week6/NamikoToriyama/model"
)

func TestListTweet(t *testing.T) {
	ctx := context.Background()
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
	// Set return value of mock.
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockInfra := mock_app.NewMockTweetRepository(ctrl)
	mockInfra.EXPECT().List(gomock.Any()).Return(want, nil)

	target := &tweetService{
		client: mockInfra,
	}

	// Execute test.
	got, err := target.ListTweet(ctx)
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("want = %v, got = %v", want, got)
	}
}

func TestPostTweet(t *testing.T) {
	ctx := context.Background()
	testTweet := &model.Tweet{
		ID:       "12345",
		Username: "foo",
		Tweet:    "hello world",
	}

	// Set return value of mock.
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockInfra := mock_app.NewMockTweetRepository(ctrl)
	returnKey := 1234
	mockInfra.EXPECT().Put(gomock.Any(), gomock.Any()).Return(returnKey, nil)

	target := &tweetService{
		client: mockInfra,
	}

	// Execute test.
	if err := target.PostTweet(ctx, testTweet); err != nil {
		t.Error(err)
	}
}

func TestGetTweet(t *testing.T) {
	ctx := context.Background()
	now := time.Now()
	want := &model.Tweet{
		ID:           "12345",
		Username:     "foo",
		Tweet:        "hello world",
		RegisterDate: now,
	}
	// Set return value of mock.
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockInfra := mock_app.NewMockTweetRepository(ctrl)
	mockInfra.EXPECT().Get(gomock.Any(), gomock.Any()).Return(want, nil)

	target := &tweetService{
		client: mockInfra,
	}

	// Execute test.
	got, err := target.GetTweet(ctx, want.ID)
	if err != nil {
		t.Error(err)
	}
	if got != want {
		t.Fatalf("want = %v, got = %v", want, got)
	}
}

func TestUpdateTweet(t *testing.T) {
	ctx := context.Background()
	now := time.Now()
	ex := &model.Tweet{
		ID:           "12345",
		Username:     "foo",
		Tweet:        "hello world",
		RegisterDate: now,
	}
	want := &model.Tweet{
		ID:           "12345",
		Username:     "foo",
		Tweet:        "Hello world!!",
		RegisterDate: now,
	}

	// Set return value of mock.
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockInfra := mock_app.NewMockTweetRepository(ctrl)
	mockInfra.EXPECT().Get(gomock.Any(), gomock.Any()).Return(ex, nil)

	returnKey := 1234
	mockInfra.EXPECT().Put(gomock.Any(), gomock.Any()).Return(returnKey, nil)

	target := &tweetService{
		client: mockInfra,
	}

	// Execute test.
	got, err := target.UpdateTweet(ctx, want)
	if err != nil {
		t.Error(err)
	}
	if got != want {
		t.Fatalf("want = %v, got = %v", want, got)
	}
}

func TestDeleteTweet(t *testing.T) {
	ctx := context.Background()
	testTweet := &model.Tweet{
		ID: "12345",
	}

	// Set return value of mock.
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockInfra := mock_app.NewMockTweetRepository(ctrl)
	mockInfra.EXPECT().Delete(gomock.Any(), gomock.Any()).Return(nil)
	target := &tweetService{
		client: mockInfra,
	}

	// Execute test.
	if err := target.DeleteTweet(ctx, testTweet.ID); err != nil {
		t.Error(err)
	}
}
