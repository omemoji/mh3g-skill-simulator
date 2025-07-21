FROM golang:1.24.5-alpine3.21

WORKDIR /go

COPY go/go.mod go/go.sum ./

RUN go mod download

COPY go/ ./

RUN go build -o main cmd/main.go

FROM alpine:3.21

WORKDIR /go

COPY --from=0 /go/main .
COPY --from=0 /go/mh3g_data/ mh3g_data/

CMD ["./main"]