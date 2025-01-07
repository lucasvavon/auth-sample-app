# Build the application from source
FROM golang:alpine3.21 AS build-stage

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o auth-sample-app ./cmd/main.go

EXPOSE 1323

ENTRYPOINT ["/app/auth-sample-app"]

# Run the tests in the container
FROM build-stage AS run-test-stage

RUN go test -v ./...

