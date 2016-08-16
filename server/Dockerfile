FROM golang

ADD . /go/src/
ENV GOBIN=/go/bin
ENV ASSET_ROOT=/go/src/assets
RUN go install /go/src/main.go

# Run the main command by default when the container starts.
ENTRYPOINT /go/bin/main

EXPOSE 3000
