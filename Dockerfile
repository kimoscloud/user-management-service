LABEL authors="sebastianbogado"
FROM golang:1.21 as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
WORKDIR /app/cmd
RUN CGO_ENABLED=0 GOOS=linux go build -o main .
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/cmd/main .
EXPOSE 10101
CMD ["./main"]
