package courses

import (
	"net/http"

	"github.com/gin-gonic/gin"
	mooc "github.com/jlezcanof/go-hexagonal_http_api-course/03-02-repository-test/internal"
)

type getResponse struct {
	ID       string `json:"id" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Duration string `json:"duration" binding:"required"`
}

// CreateHandler returns an HTTP handler for courses creation.
func GetHandler(courseRepository mooc.CourseRepository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var courses, err = courseRepository.GetAll(ctx)
		if err != nil {
			// Si quiero devolver error en ves de la lista se rompe me genera un error de unmarshal
			ctx.JSON(http.StatusInternalServerError, []getResponse{})
			return
		}
		response := make([]getResponse, 0, len(courses))
		for _, course := range courses {
			response = append(response, getResponse{
				ID:       course.ID().String(),
				Name:     course.Name().String(),
				Duration: course.Duration().String(),
			})
		}
		ctx.JSON(http.StatusOK, response)
	}
}
