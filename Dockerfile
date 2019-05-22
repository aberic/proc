FROM golang:1.12.3 as builder
LABEL app="proc" by="ennoo"
ENV REPO=$GOPATH/src/github.com/ennoo/proc
WORKDIR $REPO
RUN git clone https://github.com/ennoo/proc.git ../proc && \
 go build -o $REPO/proc $REPO/runner/proc.go
FROM centos:7
WORKDIR /root/
COPY --from=builder /go/src/github.com/ennoo/proc/proc .
EXPOSE 19637
CMD ./proc

# https://microbadger.com/labels
LABEL io.github.ennoo.name="Proc Image" \
      io.github.ennoo.description="Linux System proc api developed with golang" \
      io.github.ennoo.url="https://github.com/ennoo/proc" \
      io.github.ennoo.license="Apache License 2.0" \
      io.github.ennoo.docker.dockerfile="Dockerfile" \
      io.github.ennoo.vcs-type="Git" \
      io.github.ennoo.vcs-url="https://github.com/ennoo/proc.git" \
      io.github.ennoo.vendor="ENNOO"