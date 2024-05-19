# syntax=docker/dockerfile:1

# stage1: Build the application from source
FROM golang:1.21.5-alpine AS build-stage

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /main

# stage2: Run the tests in the container
FROM build-stage AS run-test-stage
RUN go test -v ./...

# stage3: Deploy the application binary into a lean image
FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /

COPY --from=build-stage /main /main

EXPOSE 8080

USER root:root

ENTRYPOINT ["/main"]
