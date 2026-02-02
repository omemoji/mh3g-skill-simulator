FROM golang:1.25.6-trixie as builder

WORKDIR /app

COPY go/go.mod go/go.sum ./

RUN go mod download

COPY go/ ./

RUN CGO_ENABLED=0 go build -o main cmd/main.go

FROM gcr.io/distroless/static-debian13

WORKDIR /app

COPY --from=builder /app/ ./ 

CMD ["./main"]