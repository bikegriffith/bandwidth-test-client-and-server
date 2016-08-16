FROM golang

ENV GOBIN=/go/bin
ADD . /go/src/
RUN go install /go/src/main.go

# Run the main command by default when the container starts.
ENTRYPOINT /go/bin/main

# Service listens on port 3000
EXPOSE 8080
