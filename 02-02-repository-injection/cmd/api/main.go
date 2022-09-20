package main

import (
	"log"

	"github.com/jlezcanof/go-hexagonal_http_api-course/02-02-repository-injection/cmd/api/bootstrap"
)

func main() {
	if err := bootstrap.Run(); err != nil {
		log.Fatal(err)
	}
}
