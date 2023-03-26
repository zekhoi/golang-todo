# Start from the latest golang base image
FROM golang:alpine AS builder

LABEL maintainer="Khoironi Kurnia Syah <khoironidev@gmail.com>"

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/server/main.go

# Start from a Debian image with the latest version of Go installed
FROM alpine:latest

RUN apk --no-cache add ca-certificates
RUN apk add --no-cache make

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage. Observe we also copied the .env file
COPY --from=builder /app/main .
COPY --from=builder /app/.env .

# Expose port 3030 to the outside world
EXPOSE 3030

#Command to run the executable
CMD ["./main"]
