FROM golang:alpine as builder
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn
WORKDIR /app
COPY . .
RUN go mod download
RUN go env && make build

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/app .
COPY --from=builder /app/.env* .
ENTRYPOINT /app/app
