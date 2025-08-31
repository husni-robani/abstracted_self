# Stage 1: Build
FROM golang:1.23-alpine AS build
WORKDIR /app

# Copy go.mod and go.sum first for caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the code
COPY . .

# Build the binary
RUN go build -o main ./cmd

# Stage 2: Run
FROM alpine:latest
WORKDIR /app

COPY --from=build /app/main .

# Copy profile_data.json
RUN mkdir data
COPY internal/db/profile_data.json data/.

EXPOSE 8080
CMD ["./main"]
