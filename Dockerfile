FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o main .
RUN go test -v ./... -coverprofile=coverage.out

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /app/main .

EXPOSE 8080

CMD ["./main"]