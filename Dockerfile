
FROM golang:alpine as builder

WORKDIR /go/release
#RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
RUN apk update && apk add tzdata

COPY go.mod ./go.mod
RUN go mod tidy
COPY . .
RUN pwd && ls

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -a -installsuffix cgo -o go-admin .

FROM alpine
# ENV GOPROXY https://goproxy.cn/

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories

RUN apk update --no-cache
RUN apk add --update gcc g++ libc6-compat
RUN apk add --no-cache ca-certificates
RUN apk add --no-cache tzdata
ENV TZ Asia/Shanghai

COPY --from=builder /go/release/go-admin /
COPY --from=builder /go/release/config/settings.yml /config/settings.yml
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
COPY ./go-admin-db.db /go-admin-db.db
EXPOSE 8000
CMD ["/go-admin","server","-c", "/config/settings.yml"]