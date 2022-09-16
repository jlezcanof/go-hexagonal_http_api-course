package mysql

import (
	"context"
	"database/sql"
	"fmt"

	mooc "github.com/CodelyTV/go-hexagonal_http_api-course/03-02-repository-test/internal"
	"github.com/huandu/go-sqlbuilder"
)

// CourseRepository is a MySQL mooc.CourseRepository implementation.
type CourseRepository struct {
	db *sql.DB
}

// NewCourseRepository initializes a MySQL-based implementation of mooc.CourseRepository.
func NewCourseRepository(db *sql.DB) *CourseRepository {
	return &CourseRepository{
		db: db,
	}
}

// Save implements the mooc.CourseRepository interface.
func (r *CourseRepository) Save(ctx context.Context, course mooc.Course) error {
	courseSQLStruct := sqlbuilder.NewStruct(new(sqlCourse))
	query, args := courseSQLStruct.InsertInto(sqlCourseTable, sqlCourse{
		ID:       course.ID().String(),
		Name:     course.Name().String(),
		Duration: course.Duration().String(),
	}).Build()

	_, err := r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("error trying to persist course on database: %v", err)
	}

	return nil
}

// Get all courses the mooc.CourseRepository interface.
func (r *CourseRepository) GetAll(ctx context.Context) (courses []mooc.Course, err error) {
	coursesSQLStruct := sqlbuilder.NewSelectBuilder()
	coursesSQLStruct.Select("id", "name", "duration")
	coursesSQLStruct.From(sqlCourseTable)

	query, args := coursesSQLStruct.Build()

	rows, err := r.db.QueryContext(ctx, query, args)

	if err != nil {
		return nil, fmt.Errorf("error trying to get courses on database: %v", err)
	}

	defer rows.Close()
	courses = []mooc.Course{}
	for rows.Next() {
		var sqlCourse sqlCourse
		err := rows.Scan(sqlCourse.ID, sqlCourse.Name, sqlCourse.Duration)
		if err != nil {
			return nil, err
		}
		course, err := mooc.NewCourse(sqlCourse.ID, sqlCourse.Name, sqlCourse.Duration)
		if err != nil {
			return nil, err
		}
		courses = append(courses, course)
	}

	return courses, nil
}
