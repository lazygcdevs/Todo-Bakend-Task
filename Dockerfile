# ---- Build Stage ----
FROM golang:1.23 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main .

# ---- Run Stage ----
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/startup.sh ./startup.sh
RUN chmod +x ./main && chmod +x ./startup.sh
EXPOSE 8080
CMD ["./main"] 