FROM golang:1.24.5-alpine3.21

WORKDIR /go

COPY go/go.mod go/go.sum ./

RUN go mod download

COPY go/ ./

RUN go build -o main cmd/main.go

FROM alpine:3.21

WORKDIR /go

COPY --from=0 /app/main .   

CMD ["./main"]