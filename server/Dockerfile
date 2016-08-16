FROM golang

ADD . /go/src/
ENV GOBIN=/go/bin
RUN go install /go/src/main.go

ENV STATIC_ROOT=/go/src/static
ENV MAX_UPLOAD_MB=10
ENTRYPOINT /go/bin/main

EXPOSE 3000
