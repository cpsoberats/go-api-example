// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	entities "api/src/social/artist/entities"

	mock "github.com/stretchr/testify/mock"
)

// IArtistService is an autogenerated mock type for the IArtistService type
type IArtistService struct {
	mock.Mock
}

// GetArtistList provides a mock function with given fields: from, size
func (_m *IArtistService) GetArtistList(from uint64, size uint64) []entities.Artist {
	ret := _m.Called(from, size)

	var r0 []entities.Artist
	if rf, ok := ret.Get(0).(func(uint64, uint64) []entities.Artist); ok {
		r0 = rf(from, size)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entities.Artist)
		}
	}

	return r0
}
