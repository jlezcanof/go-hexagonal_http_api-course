// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package storagemocks

import (
	context "context"

	mooc "github.com/jlezcanof/go-hexagonal_http_api-course/03-02-repository-test/internal"
	mock "github.com/stretchr/testify/mock"
)

// CourseRepository is an autogenerated mock type for the CourseRepository type
type CourseRepository struct {
	mock.Mock
}

// Save provides a mock function with given fields: ctx, course
func (_m *CourseRepository) Save(ctx context.Context, course mooc.Course) error {
	ret := _m.Called(ctx, course)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, mooc.Course) error); ok {
		r0 = rf(ctx, course)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAll provides a mock function with given fields: ctx
func (_m *CourseRepository) GetAll(ctx context.Context) ([]mooc.Course, error) {
	ret := _m.Called(ctx)

	var r0 []mooc.Course
	if rf, ok := ret.Get(0).(func(context.Context) []mooc.Course); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]mooc.Course)
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
