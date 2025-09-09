FROM golang:1.24-alpine

WORKDIR /app

COPY topz/ ./
RUN go mod download

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-gs-ping cmd/server/main.go

EXPOSE 8088

# Run
CMD ["/docker-gs-ping"]