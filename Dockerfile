FROM golang:onbuild
Maintainer  Parthiban Srinivasan 

ADD  .  /gopath/src/github.com/parthiban-srinivasan/goproj

RUN go install github.com/parthiban-srinivasan/goproj

ENTRYPOINT /gopath/bin/

EXPOSE 8080