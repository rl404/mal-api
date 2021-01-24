# Golang base image
FROM golang:1.15 as go_builder
LABEL stage="mal_api_builder"
WORKDIR /go/src/github.com/rl404/mal-api
COPY . .
WORKDIR /go/src/github.com/rl404/mal-api/cmd/malscraper
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o malscraper

# New stage from scratch
FROM alpine:3.13
RUN apk add --no-cache ca-certificates
WORKDIR /docker/bin
COPY --from=go_builder /go/src/github.com/rl404/mal-api/cmd/malscraper/malscraper malscraper
CMD ["/docker/bin/malscraper"]