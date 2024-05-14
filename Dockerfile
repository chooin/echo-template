FROM golang:alpine as builder
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn
ENV TZ=Asia/Shanghai
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone
WORKDIR /app
COPY . .
RUN go mod download
RUN go env
RUN go build -o app .
EXPOSE 8080
ENTRYPOINT /app/app
