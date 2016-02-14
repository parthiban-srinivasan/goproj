FROM scratch
MAINTAINER  Parthiban Srinivasan <parthiban.srinivasan@gmail.com> 

ADD  testmicro testmicro

ENTRYPOINT ["/testmicro"]

EXPOSE 8080
