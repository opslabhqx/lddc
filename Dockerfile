# syntax=docker/dockerfile:1

# renovate: datasource=docker depName=alpine
ARG ALPINE_VERSION=3.20.1
# renovate: datasource=docker depName=golang
ARG GO_VERSION=1.22
# renovate: datasource=github-releases depName=opslabhqx/lddc
ARG LDDC_VERSION=1.0.0

FROM alpine:${ALPINE_VERSION} AS base

FROM golang:${GO_VERSION}-alpine AS build

WORKDIR /app

COPY go.mod .

RUN go mod download

COPY . .

RUN go mod tidy && CGO_ENABLED="0" go build -o lddc app/*.go

FROM base

COPY --from=build /app/lddc /usr/local/bin
