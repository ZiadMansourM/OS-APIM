FROM golang:1.18.2-alpine3.16 AS builder

WORKDIR /app
COPY go.mod ./
RUN go mod download
ADD . .
RUN CGO_ENABLED=0 go build -o server.o

FROM scratch
WORKDIR /app
COPY --from=builder /app/server.o .

EXPOSE 8080

ENTRYPOINT ["./server.o"]