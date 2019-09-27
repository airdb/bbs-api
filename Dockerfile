# Stage 1: Builder
FROM golang
MAINTAINER info@airdb.com

ENV PROJ=bbs-api
ENV BIN=main
ENV GITHUB=github.com/airdb/${PROJ}
ENV BUILDDIR=/go/src/${GITHUB}
ENV DEPLOYDIR=/srv/${PROJ}
ENV RUNBIN=${DEPLOYDIR}/${BIN}

WORKDIR ${BUILDDIR}

ADD . ${BUILDDIR}

RUN go get -u github.com/swaggo/swag/cmd/swag && \
	GO111MODULE=off swag init -g web/router.go && \
	go build -o main main.go

# Stage 2: Release the binary from the builder stage
FROM ubuntu

ENV PROJ=bbs-api
ENV BIN=main
ENV GITHUB=github.com/airdb/${PROJ}
ENV BUILDDIR=/go/src/${GITHUB}
ENV DEPLOYDIR=/srv/${PROJ}
ENV RUNBIN=${DEPLOYDIR}/${BIN}

COPY --from=0 ${BUILDDIR}/ ${DEPLOYDIR}

EXPOSE 8080

WORKDIR ${DEPLOYDIR}
CMD [${RUNBIN}]
