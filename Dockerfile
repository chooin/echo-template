FROM golang as builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build

FROM ubuntu
WORKDIR /app
COPY --from=builder /app/app /app/app
RUN apt update
RUN apt install -y --no-install-recommends ca-certificates curl
EXPOSE 8080
