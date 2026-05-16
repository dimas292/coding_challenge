FROM golang:1.26-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN mkdir -p /app/bin \
  && CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /app/bin/server . \
  && chmod 0755 /app/bin/server

FROM alpine:3.21

WORKDIR /app

COPY --from=build /app/bin/server /app/server
COPY --from=build /app/config.yml /app/config.yml
RUN chmod 0755 /app/server

EXPOSE 4444

ENTRYPOINT ["/app/server"]
