FROM golang:1.25-alpine AS builder
RUN apk add  make gcc musl-dev
WORKDIR /app
COPY . .
RUN make build
FROM alpine:latest
RUN apk add ca-certificates
WORKDIR /app
COPY --from=builder /app/bin/pasteBin .
EXPOSE 8080
CMD ["./pasteBin"]
