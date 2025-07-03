# base image
FROM golang:1.23-alpine as builder

RUN mkdir /app

COPY . /app

WORKDIR /app

    # build the go program
RUN CGO_ENABLED=0 go build -o abstracted_self ./cmd

RUN chmod +x /app/abstracted_self

# build a tiny docker image
FROM alpine:latest

RUN mkdir /app

COPY --from=builder /app/abstracted_self /app

CMD [ "app/abstracted_self" ]