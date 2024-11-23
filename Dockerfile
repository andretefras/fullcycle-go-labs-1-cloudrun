FROM golang:1.23 as builder
WORKDIR /app
COPY . .
RUN go mod download && \
    go mod verify && \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build --ldflags="-w -s" -o app ./cmd/main.go \
FROM scratch
WORKDIR /app
COPY --from=builder /app/app .
ENTRYPOINT ["./app"]