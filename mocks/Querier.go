// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	context "context"

	models "github.com/duyledat197/interview-hao/internal/models"
	mock "github.com/stretchr/testify/mock"
)

// Querier is an autogenerated mock type for the Querier type
type Querier struct {
	mock.Mock
}

// GetHub provides a mock function with given fields: ctx, id
func (_m *Querier) GetHub(ctx context.Context, id int64) (models.Hub, error) {
	ret := _m.Called(ctx, id)

	var r0 models.Hub
	if rf, ok := ret.Get(0).(func(context.Context, int64) models.Hub); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(models.Hub)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUser provides a mock function with given fields: ctx, id
func (_m *Querier) GetUser(ctx context.Context, id int64) (models.User, error) {
	ret := _m.Called(ctx, id)

	var r0 models.User
	if rf, ok := ret.Get(0).(func(context.Context, int64) models.User); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(models.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserHub provides a mock function with given fields: ctx
func (_m *Querier) GetUserHub(ctx context.Context) ([]models.GetUserHubRow, error) {
	ret := _m.Called(ctx)

	var r0 []models.GetUserHubRow
	if rf, ok := ret.Get(0).(func(context.Context) []models.GetUserHubRow); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.GetUserHubRow)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewQuerier interface {
	mock.TestingT
	Cleanup(func())
}

// NewQuerier creates a new instance of Querier. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewQuerier(t mockConstructorTestingTNewQuerier) *Querier {
	mock := &Querier{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}