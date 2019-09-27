# Stage 1: Builder
FROM golang
MAINTAINER info@airdb.com

WORKDIR  /go/src/github.com/airdb/dns-api

ADD . $WORKDIR

RUN go get -u github.com/swaggo/swag/cmd/swag && \
	GO111MODULE=off swag init -g web/router.go && \
	go build -o dns-api main.go

# Stage 2: Copy the binary from the builder stage
FROM ubuntu
COPY --from=0 /go/src/github.com/airdb/dns-api /srv/

EXPOSE 8080

WORKDIR  /srv
CMD ["/srv/dns-api"]
