// Code generated by MockGen. DO NOT EDIT.
// Source: tweet_service.go

// Package mock_app is a generated GoMock package.
package mock_app

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	model "github.com/mercari/Week6/Week6/NamikoToriyama/model"
)

// MockTweetService is a mock of TweetService interface
type MockTweetService struct {
	ctrl     *gomock.Controller
	recorder *MockTweetServiceMockRecorder
}

// MockTweetServiceMockRecorder is the mock recorder for MockTweetService
type MockTweetServiceMockRecorder struct {
	mock *MockTweetService
}

// NewMockTweetService creates a new mock instance
func NewMockTweetService(ctrl *gomock.Controller) *MockTweetService {
	mock := &MockTweetService{ctrl: ctrl}
	mock.recorder = &MockTweetServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTweetService) EXPECT() *MockTweetServiceMockRecorder {
	return m.recorder
}

// ListTweet mocks base method
func (m *MockTweetService) ListTweet(arg0 context.Context) ([]*model.Tweet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTweet", arg0)
	ret0, _ := ret[0].([]*model.Tweet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTweet indicates an expected call of ListTweet
func (mr *MockTweetServiceMockRecorder) ListTweet(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTweet", reflect.TypeOf((*MockTweetService)(nil).ListTweet), arg0)
}

// PostTweet mocks base method
func (m *MockTweetService) PostTweet(arg0 context.Context, arg1 *model.Tweet) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostTweet", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// PostTweet indicates an expected call of PostTweet
func (mr *MockTweetServiceMockRecorder) PostTweet(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostTweet", reflect.TypeOf((*MockTweetService)(nil).PostTweet), arg0, arg1)
}

// GetTweet mocks base method
func (m *MockTweetService) GetTweet(arg0 context.Context, arg1 string) (*model.Tweet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTweet", arg0, arg1)
	ret0, _ := ret[0].(*model.Tweet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTweet indicates an expected call of GetTweet
func (mr *MockTweetServiceMockRecorder) GetTweet(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTweet", reflect.TypeOf((*MockTweetService)(nil).GetTweet), arg0, arg1)
}

// UpdateTweet mocks base method
func (m *MockTweetService) UpdateTweet(arg0 context.Context, arg1 *model.Tweet) (*model.Tweet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTweet", arg0, arg1)
	ret0, _ := ret[0].(*model.Tweet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateTweet indicates an expected call of UpdateTweet
func (mr *MockTweetServiceMockRecorder) UpdateTweet(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTweet", reflect.TypeOf((*MockTweetService)(nil).UpdateTweet), arg0, arg1)
}

// DeleteTweet mocks base method
func (m *MockTweetService) DeleteTweet(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTweet", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTweet indicates an expected call of DeleteTweet
func (mr *MockTweetServiceMockRecorder) DeleteTweet(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTweet", reflect.TypeOf((*MockTweetService)(nil).DeleteTweet), arg0, arg1)
}
