# Stage 1: Builder
FROM golang
MAINTAINER info@airdb.com

ENV PROJ=bbs-sync
ENV BIN=main
ENV GITHUB=github.com/airdb/${PROJ}
ENV BUILDDIR=/go/src/${GITHUB}
ENV DEPLOYDIR=/srv/${PROJ}
ENV RUNBIN=${DEPLOYDIR}/${BIN}

WORKDIR ${BUILDDIR}

ADD . ${BUILDDIR}

RUN go build -o main cmd/bbs_sync.go

# Stage 2: Release the binary from the builder stage
FROM ubuntu

ENV PROJ=bbs-sync
ENV BIN=main
ENV GITHUB=github.com/airdb/${PROJ}
ENV BUILDDIR=/go/src/${GITHUB}
ENV DEPLOYDIR=/srv/${PROJ}
ENV RUNBIN=${DEPLOYDIR}/${BIN}

COPY --from=0 ${BUILDDIR}/ ${DEPLOYDIR}

EXPOSE 8080

WORKDIR ${DEPLOYDIR}
CMD ${RUNBIN}
