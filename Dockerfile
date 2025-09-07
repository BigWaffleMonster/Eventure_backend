FROM golang:1.23 AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -ldflags="-w -s -extldflags '-static'" \
    -trimpath \
    -o /app/cmd/main ./cmd/


# Stage 2: Create a minimal runtime image
FROM alpine:3.21
WORKDIR /app

# Copy binary from builder stage (в ту же директорию)
COPY --from=builder /app/cmd/main .

# Make binary executable
RUN chmod +x main

# Expose port
EXPOSE ${APP_PORT}

# Run the application
CMD ["./main"]
