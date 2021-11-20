FROM golang:1.17-alpine AS builder

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

WORKDIR /build

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o ./dist/main ./cmd

FROM alpine

ENV APP_ENV=development
ENV APP_PORT=80

WORKDIR /build

COPY --from=builder /build/dist/main ./
COPY --from=builder /build/keys ./keys

ENTRYPOINT [ "/build/main" ]
