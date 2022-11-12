package bootstrap

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jlezcanof/go-hexagonal_http_api-course/02-01-post-course-endpoint/internal/platform/server"
)

const (
	host = "localhost"
	port = 8080
)

func Run() error {
	srv := server.New(host, port)
	return srv.Run()
}
