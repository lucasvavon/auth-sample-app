# Build the application from source
FROM golang:1.23-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o auth-sample-app ./cmd/main.go

EXPOSE 1323

ENTRYPOINT ["/app/auth-sample-app"]