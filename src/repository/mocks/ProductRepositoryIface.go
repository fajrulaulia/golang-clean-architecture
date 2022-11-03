// Code generated by mockery v2.14.1. DO NOT EDIT.

package mocks

import (
	model "github.com/fajrulaulia/arsitektur-bersih/src/model"
	mock "github.com/stretchr/testify/mock"
)

// ProductRepositoryIface is an autogenerated mock type for the ProductRepositoryIface type
type ProductRepositoryIface struct {
	mock.Mock
}

// Create provides a mock function with given fields: code, name, price
func (_m *ProductRepositoryIface) Create(code string, name string, price float64) error {
	ret := _m.Called(code, name, price)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, float64) error); ok {
		r0 = rf(code, name, price)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetByID provides a mock function with given fields: id
func (_m *ProductRepositoryIface) GetByID(id string) (*model.Product, error) {
	ret := _m.Called(id)

	var r0 *model.Product
	if rf, ok := ret.Get(0).(func(string) *model.Product); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Product)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewProductRepositoryIface interface {
	mock.TestingT
	Cleanup(func())
}

// NewProductRepositoryIface creates a new instance of ProductRepositoryIface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewProductRepositoryIface(t mockConstructorTestingTNewProductRepositoryIface) *ProductRepositoryIface {
	mock := &ProductRepositoryIface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}