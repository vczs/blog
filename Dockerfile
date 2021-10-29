FROM golang

LABEL maintainer="vczs"

WORKDIR $GOPATH/src/godocker

ADD . $GOPATH/src/godocker

RUN go build main.go

EXPOSE 80

ENTRYPOINT ["./main"]