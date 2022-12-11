FROM golang:1.19 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN CGO_ENABLED=0 go build -v ./cmd/flp

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/flp ./

RUN ls -al

CMD ["./flp"]
