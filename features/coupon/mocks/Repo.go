// Code generated by mockery v2.37.1. DO NOT EDIT.

package mocks

import (
	coupon "user_barang/features/coupon"

	mock "github.com/stretchr/testify/mock"
)

// Repo is an autogenerated mock type for the Repo type
type Repo struct {
	mock.Mock
}

// GetCoupon provides a mock function with given fields: userId, page, limit
func (_m *Repo) GetCoupon(userId uint, page int, limit int) ([]coupon.Coupon, error) {
	ret := _m.Called(userId, page, limit)

	var r0 []coupon.Coupon
	var r1 error
	if rf, ok := ret.Get(0).(func(uint, int, int) ([]coupon.Coupon, error)); ok {
		return rf(userId, page, limit)
	}
	if rf, ok := ret.Get(0).(func(uint, int, int) []coupon.Coupon); ok {
		r0 = rf(userId, page, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]coupon.Coupon)
		}
	}

	if rf, ok := ret.Get(1).(func(uint, int, int) error); ok {
		r1 = rf(userId, page, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCouponTanpa provides a mock function with given fields: page, limit
func (_m *Repo) GetCouponTanpa(page int, limit int) ([]coupon.Coupon, error) {
	ret := _m.Called(page, limit)

	var r0 []coupon.Coupon
	var r1 error
	if rf, ok := ret.Get(0).(func(int, int) ([]coupon.Coupon, error)); ok {
		return rf(page, limit)
	}
	if rf, ok := ret.Get(0).(func(int, int) []coupon.Coupon); ok {
		r0 = rf(page, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]coupon.Coupon)
		}
	}

	if rf, ok := ret.Get(1).(func(int, int) error); ok {
		r1 = rf(page, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertCoupon provides a mock function with given fields: userId, newCoupon
func (_m *Repo) InsertCoupon(userId uint, newCoupon coupon.Coupon) (coupon.Coupon, error) {
	ret := _m.Called(userId, newCoupon)

	var r0 coupon.Coupon
	var r1 error
	if rf, ok := ret.Get(0).(func(uint, coupon.Coupon) (coupon.Coupon, error)); ok {
		return rf(userId, newCoupon)
	}
	if rf, ok := ret.Get(0).(func(uint, coupon.Coupon) coupon.Coupon); ok {
		r0 = rf(userId, newCoupon)
	} else {
		r0 = ret.Get(0).(coupon.Coupon)
	}

	if rf, ok := ret.Get(1).(func(uint, coupon.Coupon) error); ok {
		r1 = rf(userId, newCoupon)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewRepo creates a new instance of Repo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRepo(t interface {
	mock.TestingT
	Cleanup(func())
}) *Repo {
	mock := &Repo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
