FROM golang:alpine AS build

RUN apk add --update git
WORKDIR /go/src/github.com/jlezcanof/go-hexagonal_http_api-course
COPY . .
RUN CGO_ENABLED=0 go build -o /go/bin/jlezcanof-mooc-api 08-03-debugging/cmd/api/main.go

# Building image with the binary
FROM scratch
COPY --from=build /go/bin/jlezcanof-mooc-api /go/bin/jlezcanof-mooc-api
ENTRYPOINT ["/go/bin/jlezcanof-mooc-api"]
