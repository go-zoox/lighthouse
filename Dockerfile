# Builder
FROM golang:1.19-alpine as builder

RUN         apk add --no-cache gcc g++ make git

WORKDIR     /app

COPY        go.mod .

COPY        go.sum .

RUN         go mod download

# ARG         VERSION=unknown

# ARG         BUILD_TIME=unknown

# ARG         COMMIT_HASH=unknown

COPY        . .

# RUN         CGO_ENABLED=1 \
#             GOOS=linux \
#             GOARCH=amd64 \
#             go build \
#               -trimpath \
#               -ldflags '\
#                 -X "github.com/go-zoox/lighthouse/constants.Version=${VERSION}" \
#                 -X "github.com/go-zoox/lighthouse/constants.BuildTime=${BUILD_TIME}" \
#                 -X "github.com/go-zoox/lighthouse/constants.CommitHash=${COMMIT_HASH}" \
#                 -w -s -buildid= \
#               ' \
#               -v -o lighthouse

RUN GOOS=linux \
  GOARCH=amd64 \
  go build \
  -trimpath \
  -ldflags '-w -s -buildid=' \
  -v -o lighthouse

# Product
FROM        golang:1.19-alpine

LABEL       MAINTAINER="Zero<tobewhatwewant@gmail.com>"

LABEL       org.opencontainers.image.source="https://github.com/go-zoox/lighthouse"

ARG         VERSION=v1.0.0

ENV         MODE=production

COPY        --from=builder /app/lighthouse /bin

COPY        conf/lighthouse.yaml /conf/lighthouse.yaml

ENV         VERSION=${VERSION}

CMD  ["lighthouse", "-c", "/conf/lighthouse.yaml"]
