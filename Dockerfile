ARG GO_VERSION=1.24

FROM golang:$GO_VERSION-alpine AS builder

WORKDIR /workspace

COPY go.mod go.sum ./
RUN go mod download
RUN go install github.com/a-h/templ/cmd/templ@latest

COPY . .

RUN templ generate

RUN CGO_ENABLED=0 GOOS=linux go build -o /app .

# Run the tests in the container
FROM builder AS tester
RUN go test -v ./...

# Deploy the application binary into a lean image
FROM alpine:latest AS releaser

WORKDIR /

COPY --from=builder /app /app

RUN addgroup -S nonroot && adduser -S nonroot -G nonroot

USER nonroot:nonroot

ENTRYPOINT ["/app"]
