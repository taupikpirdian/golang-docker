# Gunakan image Go resmi
FROM golang:1.21-alpine

# Set environment
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Set working directory
WORKDIR /app

# Copy go.mod dan go.sum untuk cache
COPY go.mod ./

# Unduh dependencies
RUN go mod download

# Copy semua file ke dalam container
COPY src ./src

# Build aplikasi
RUN go build -o main ./src

# Expose port 8081 untuk HTTP server
EXPOSE 8081

# Jalankan aplikasi
CMD ["./main"]
