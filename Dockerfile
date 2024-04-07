FROM golang:1.20.1-alpine AS builder

ENV GOOS=linux GOARCH=amd64

WORKDIR /books-recommender

COPY . .

RUN go build -o books-recommender.bin

FROM alpine AS final

# Copying the binary
COPY --from=builder /books-recommender/books-recommender.bin /books-recommender/books-recommender.bin

CMD ["./books-recommender/books-recommender.bin"]
