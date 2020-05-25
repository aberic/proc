FROM golang:1.14.3 as builder
LABEL app="proc" by="aberic"
ENV REPO=$GOPATH/src/github.com/aberic/proc
WORKDIR $REPO
ADD . $REPO
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build
FROM centos:7
WORKDIR /root/
COPY --from=builder /go/src/github.com/aberic/proc/proc .
EXPOSE 19637
CMD ./proc