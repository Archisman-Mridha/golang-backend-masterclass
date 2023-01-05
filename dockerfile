#* building stage
FROM golang:1.19-alpine AS builder

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o build

#* packaging stage
FROM alpine:latest AS packager

WORKDIR /

COPY --from=builder /app/build .

EXPOSE 4000

CMD [ "./build" ]