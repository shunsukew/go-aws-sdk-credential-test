# syntax=docker/dockerfile:1

FROM golang:1.22 AS builder

WORKDIR /go/src/github.com/shunsukew/go-kms-credential-test

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/main ./main.go

# ====================

FROM alpine:latest

RUN apk --no-cache add ca-certificates
RUN apk add libc6-compat

WORKDIR /root/

COPY --from=builder /go/src/github.com/shunsukew/go-kms-credential-test/bin/main ./

CMD ["./main"]
