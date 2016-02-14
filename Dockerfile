FROM scratch
MAINTAINER  Parthiban Srinivasan <parthiban.srinivasan@gmail.com> 

ADD  miservice miservice

ENTRYPOINT ["/miservice"]

EXPOSE 8080