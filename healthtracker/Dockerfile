FROM golang:1.20 as builder
ENV CGO_ENABLED=0 GOOS=linux
WORKDIR /usr/src/app

COPY . .
RUN go mod download

WORKDIR /usr/src/app/cmd/main
RUN go build -o /usr/local/bin/main -buildvcs=false

FROM alpine:latest

COPY --from=builder /usr/local/bin/main ./
CMD ["./main"]
EXPOSE 8080