FROM golang:1.20-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go

# Run stage
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .
COPY /database ./database
COPY /model/migrate ./model/migrate
COPY .env .

EXPOSE 4000

CMD /app/main
