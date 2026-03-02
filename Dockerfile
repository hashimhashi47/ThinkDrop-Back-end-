# ---------- Build Stage ----------
FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main ./cmd/api

# ---------- Run Stage ----------
FROM alpine:latest

# Install timezone data
RUN apk add --no-cache tzdata

WORKDIR /root/

COPY --from=builder /app/main .

EXPOSE 8000

CMD ["./main"]

# FROM golang:1.25-alpine

# WORKDIR /app

# RUN go install github.com/air-verse/air@latest

# ENV PATH="/go/bin:${PATH}"

# COPY go.mod go.sum ./
# RUN go mod download

# EXPOSE 8000

# CMD ["air"]
