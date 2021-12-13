FROM golang:alpine as builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/app /app/app
COPY --from=builder /app/.env* /app/.env*
EXPOSE 8080
CMD '/app/app'
