# syntax=docker/dockerfile:1

# ----------------------------------------------------------------------
# Building environment
# ----------------------------------------------------------------------

FROM golang:1.16-alpine AS go-builder
WORKDIR /app

RUN apk add build-base

COPY go.mod ./
COPY go.sum ./

COPY cmd ./cmd
COPY internal ./internal

RUN go mod tidy
RUN go build -o build/portal ./cmd/portal/portal.go

# ----------------------------------------------------------------------
# Running environment
# ----------------------------------------------------------------------

FROM alpine:latest
WORKDIR /app
COPY --from=go-builder /app /app

CMD [ "build/portal" , "start"]
