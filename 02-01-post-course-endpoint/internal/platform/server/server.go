package server

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jlezcanof/go-hexagonal_http_api-course/02-01-post-course-endpoint/internal/platform/server/handler/courses"
	"github.com/jlezcanof/go-hexagonal_http_api-course/02-01-post-course-endpoint/internal/platform/server/handler/health"
)

type Server struct {
	httpAddr string
	engine   *gin.Engine
}

func New(host string, port uint) Server {
	srv := Server{
		engine:   gin.New(),
		httpAddr: fmt.Sprintf("%s:%d", host, port),
	}

	srv.registerRoutes()
	return srv
}

func (s *Server) Run() error {
	log.Println("Server running on", s.httpAddr)
	return s.engine.Run(s.httpAddr)
}

func (s *Server) registerRoutes() {
	s.engine.GET("/health", health.CheckHandler())
	s.engine.POST("/courses", courses.CreateHandler())
}
