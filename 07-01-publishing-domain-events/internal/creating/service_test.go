package creating

import (
	"context"
	"errors"
	"testing"

	"github.com/jlezcanof/go-hexagonal_http_api-course/07-01-publishing-domain-events/internal/platform/storage/storagemocks"
	"github.com/jlezcanof/go-hexagonal_http_api-course/07-01-publishing-domain-events/kit/event/eventmocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_CourseService_CreateCourse_RepositoryError(t *testing.T) {
	courseID := "37a0f027-15e6-47cc-a5d2-64183281087e"
	courseName := "Test Course"
	courseDuration := "10 months"

	courseRepositoryMock := new(storagemocks.CourseRepository)
	courseRepositoryMock.On("Save", mock.Anything, mock.AnythingOfType("mooc.Course")).Return(errors.New("something unexpected happened"))

	eventBusMock := new(eventmocks.Bus)

	courseService := NewCourseService(courseRepositoryMock, eventBusMock)

	err := courseService.CreateCourse(context.Background(), courseID, courseName, courseDuration)

	courseRepositoryMock.AssertExpectations(t)
	eventBusMock.AssertExpectations(t)
	assert.Error(t, err)
}

func Test_CourseService_CreateCourse_EventsBusError(t *testing.T) {
	courseID := "37a0f027-15e6-47cc-a5d2-64183281087e"
	courseName := "Test Course"
	courseDuration := "10 months"

	courseRepositoryMock := new(storagemocks.CourseRepository)
	courseRepositoryMock.On("Save", mock.Anything, mock.AnythingOfType("mooc.Course")).Return(nil)

	eventBusMock := new(eventmocks.Bus)
	eventBusMock.On("Publish", mock.Anything, mock.AnythingOfType("[]event.Event")).Return(errors.New("something unexpected happened"))

	courseService := NewCourseService(courseRepositoryMock, eventBusMock)

	err := courseService.CreateCourse(context.Background(), courseID, courseName, courseDuration)

	courseRepositoryMock.AssertExpectations(t)
	eventBusMock.AssertExpectations(t)
	assert.Error(t, err)
}

func Test_CourseService_CreateCourse_Succeed(t *testing.T) {
	courseID := "37a0f027-15e6-47cc-a5d2-64183281087e"
	courseName := "Test Course"
	courseDuration := "10 months"

	courseRepositoryMock := new(storagemocks.CourseRepository)
	courseRepositoryMock.On("Save", mock.Anything, mock.AnythingOfType("mooc.Course")).Return(nil)

	eventBusMock := new(eventmocks.Bus)
	eventBusMock.On("Publish", mock.Anything, mock.AnythingOfType("[]event.Event")).Return(nil)

	courseService := NewCourseService(courseRepositoryMock, eventBusMock)

	err := courseService.CreateCourse(context.Background(), courseID, courseName, courseDuration)

	courseRepositoryMock.AssertExpectations(t)
	eventBusMock.AssertExpectations(t)
	assert.NoError(t, err)
}
