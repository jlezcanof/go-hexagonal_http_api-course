package creating

import (
	"context"
	"errors"

	mooc "github.com/jlezcanof/go-hexagonal_http_api-course/07-02-domain-events-subscriber/internal"
	"github.com/jlezcanof/go-hexagonal_http_api-course/07-02-domain-events-subscriber/internal/increasing"
	"github.com/jlezcanof/go-hexagonal_http_api-course/07-02-domain-events-subscriber/kit/event"
)

type IncreaseCoursesCounterOnCourseCreated struct {
	increasingService increasing.CourseCounterService
}

func NewIncreaseCoursesCounterOnCourseCreated(increaserService increasing.CourseCounterService) IncreaseCoursesCounterOnCourseCreated {
	return IncreaseCoursesCounterOnCourseCreated{
		increasingService: increaserService,
	}
}

func (e IncreaseCoursesCounterOnCourseCreated) Handle(_ context.Context, evt event.Event) error {
	courseCreatedEvt, ok := evt.(mooc.CourseCreatedEvent)
	if !ok {
		return errors.New("unexpected event")
	}

	return e.increasingService.Increase(courseCreatedEvt.ID())
}
