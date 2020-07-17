// Code generated by MockGen. DO NOT EDIT.
// Source: usecase.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	model "github.com/confus1on/UKM/internal/model"
	utils "github.com/confus1on/UKM/internal/utils"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockUsecaseProfile is a mock of UsecaseProfile interface
type MockUsecaseProfile struct {
	ctrl     *gomock.Controller
	recorder *MockUsecaseProfileMockRecorder
}

// MockUsecaseProfileMockRecorder is the mock recorder for MockUsecaseProfile
type MockUsecaseProfileMockRecorder struct {
	mock *MockUsecaseProfile
}

// NewMockUsecaseProfile creates a new mock instance
func NewMockUsecaseProfile(ctrl *gomock.Controller) *MockUsecaseProfile {
	mock := &MockUsecaseProfile{ctrl: ctrl}
	mock.recorder = &MockUsecaseProfileMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUsecaseProfile) EXPECT() *MockUsecaseProfileMockRecorder {
	return m.recorder
}

// GetDetailProfile mocks base method
func (m *MockUsecaseProfile) GetDetailProfile(ctx context.Context, email string) (*model.ResponseGetProfile, *utils.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDetailProfile", ctx, email)
	ret0, _ := ret[0].(*model.ResponseGetProfile)
	ret1, _ := ret[1].(*utils.Error)
	return ret0, ret1
}

// GetDetailProfile indicates an expected call of GetDetailProfile
func (mr *MockUsecaseProfileMockRecorder) GetDetailProfile(ctx, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDetailProfile", reflect.TypeOf((*MockUsecaseProfile)(nil).GetDetailProfile), ctx, email)
}

// UpdateProfile mocks base method
func (m *MockUsecaseProfile) UpdateProfile(ctx context.Context, email string, input model.InputUpdateProfile) (*model.ResponseUpdateProfile, *utils.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProfile", ctx, email, input)
	ret0, _ := ret[0].(*model.ResponseUpdateProfile)
	ret1, _ := ret[1].(*utils.Error)
	return ret0, ret1
}

// UpdateProfile indicates an expected call of UpdateProfile
func (mr *MockUsecaseProfileMockRecorder) UpdateProfile(ctx, email, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProfile", reflect.TypeOf((*MockUsecaseProfile)(nil).UpdateProfile), ctx, email, input)
}