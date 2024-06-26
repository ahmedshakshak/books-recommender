// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	entities "github.com/ahmedshakshak/books-recommender/entities"
	mock "github.com/stretchr/testify/mock"
)

// Book is an autogenerated mock type for the Book type
type Book struct {
	mock.Mock
}

// GetTheMostReadBooks provides a mock function with given fields: numOfBooks
func (_m *Book) GetTheMostReadBooks(numOfBooks int64) ([]*entities.Book, error) {
	ret := _m.Called(numOfBooks)

	var r0 []*entities.Book
	var r1 error
	if rf, ok := ret.Get(0).(func(int64) ([]*entities.Book, error)); ok {
		return rf(numOfBooks)
	}
	if rf, ok := ret.Get(0).(func(int64) []*entities.Book); ok {
		r0 = rf(numOfBooks)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entities.Book)
		}
	}

	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(numOfBooks)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewBook interface {
	mock.TestingT
	Cleanup(func())
}

// NewBook creates a new instance of Book. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewBook(t mockConstructorTestingTNewBook) *Book {
	mock := &Book{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
