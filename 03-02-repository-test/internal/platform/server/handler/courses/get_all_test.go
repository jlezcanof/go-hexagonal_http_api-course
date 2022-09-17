package courses

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	mooc "github.com/CodelyTV/go-hexagonal_http_api-course/03-02-repository-test/internal"
	"github.com/CodelyTV/go-hexagonal_http_api-course/03-02-repository-test/internal/platform/storage/storagemocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestHandler_GetAll(t *testing.T) {

	gin.SetMode(gin.TestMode)
	r := gin.New()

	t.Run("obtain empty array of courses when there are no courses", func(t *testing.T) {
		var emptyCourses []mooc.Course

		courseRepository := new(storagemocks.CourseRepository)
		courseRepository.On("GetAll", mock.Anything).Return(emptyCourses, nil)

		r.GET("/courses", GetHandler(courseRepository))

		req, err := http.NewRequest(http.MethodGet, "/courses", nil)
		//ultimo parametro
		require.NoError(t, err)

		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		res := rec.Result()
		defer res.Body.Close()

		assert.Equal(t, http.StatusOK, res.StatusCode)

		var response []mooc.Course
		if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
			log.Fatalln(err)
		}

		assert.Equal(t, []mooc.Course{}, response)

	})
}
