FROM golang:1.14.3 as builder
LABEL app="proc" by="aberic"
ENV GOPROXY=https://goproxy.io
ENV GO111MODULE=on
ENV REPO=$GOPATH/src/github.com/aberic/proc
WORKDIR $REPO
ADD . $REPO
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o $REPO/runner/proc $REPO/runner/proc.go
FROM docker.io/alpine:latest
RUN echo "https://mirror.tuna.tsinghua.edu.cn/alpine/v3.4/main" > /etc/apk/repositories && \
    apk add --update curl bash && \
    rm -rf /var/cache/apk/* && \
    mkdir -p /home
WORKDIR /root
COPY --from=builder /go/src/github.com/aberic/proc/runner/proc .
EXPOSE 19637
CMD ./proc