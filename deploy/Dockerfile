# deploy/docker/Dockerfile
FROM golang:1.19-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o skill-marketplace ./main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/skill-marketplace .
EXPOSE 8080
CMD ["./skill-marketplace"]