FROM golang:1.25-alpine AS builder
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o /out/farming-notifications-server ./...

FROM alpine:3.21
RUN adduser -D -u 1000 app
COPY --from=builder /out/farming-notifications-server /usr/local/bin/
WORKDIR /app
USER app
EXPOSE 8080
ENTRYPOINT ["farming-notifications-server"]
