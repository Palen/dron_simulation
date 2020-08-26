FROM golang:1.14-alpine AS base-image

LABEL maintainer="@Palen"
# Package dependencies
RUN apk add --no-cache --no-progress ca-certificates \
    bash \
    gcc \
    git \
    make \
    musl-dev \
    curl \
    tar \
    tzdata \
    && rm -rf /var/cache/apk/*

WORKDIR /go/src/github.com/Palen/drone_simulation
ENV GOPATH=/go/

# Dowload goimports
RUN go get golang.org/x/tools/cmd/goimports
# Download golangci-lint binary to bin folder in $GOPATH
RUN curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | bash -s -- -b $GOPATH/bin v1.27.0

ENV GO111MODULE on
COPY go.mod go.sum ./
RUN go mod download
COPY Makefile ./
COPY . .


FROM base-image as maker
EXPOSE 8080
ARG MAKE_TARGET=build
RUN make ${MAKE_TARGET}

FROM alpine:3.10

RUN addgroup -g 1000 -S app && \
    adduser -u 1000 -S app -G app

COPY --from=maker /go/src/github.com/Palen/drone_simulation/dist/drone_simulation /opt/bin/drone_simulation

CMD ["/opt/bin/drone_simulation"]
