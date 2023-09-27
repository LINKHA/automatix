FROM ubuntu:20.04
COPY main /main
WORKDIR /
EXPOSE  8999

ENTRYPOINT [ "/main" ]