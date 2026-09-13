# ====================================DEVELOPMENT====================================
FROM golang:1.23-alpine3.21 AS development
WORKDIR /app

RUN apk add --no-cache wget \
  && adduser -D -u 10001 appuser \
  && mkdir -p data storage/documents storage/images storage/icons \
  && chown -R appuser /app

USER appuser

# profile_data.json is place where profile data stored instead of using database
COPY --chown=appuser internal/db/profile_data.docker.json data/profile_data.json

COPY . .
RUN go mod download

# install air for hot reload
RUN go install github.com/air-verse/air@latest
ENV PATH="/home/appuser/go/bin:$PATH"

EXPOSE 8080
HEALTHCHECK --interval=15s --timeout=3s --retries=5 --start-period=20s CMD wget --no-verbose --tries=1 --spider http://127.0.0.1:8080/healthz || exit 1

CMD ["air"]



# ====================================PRODUCTION====================================
# Stage 1: Build
FROM golang:1.23-alpine3.21 AS build
WORKDIR /app

# Copy go.mod and go.sum first for caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the code
COPY . .

# Build statically linked binary
RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" -o /app/main ./cmd

# Stage 2: Run
FROM alpine:3.21.3 AS production
WORKDIR /app

RUN apk add --no-cache wget \
  && adduser -D -u 10001 appuser \
  && mkdir -p data storage/documents storage/images storage/icons \
  && chown -R appuser /app

COPY --from=build --chown=appuser /app/main .

# Copy profile_data.json
COPY --chown=appuser internal/db/profile_data.docker.json data/profile_data.json

USER appuser

HEALTHCHECK --interval=15s --timeout=3s --retries=5 --start-period=20s CMD wget --no-verbose --tries=1 --spider http://127.0.0.1:8080/healthz || exit 1
CMD ["./main"]
