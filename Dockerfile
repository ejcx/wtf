FROM debian
MAINTAINER Evan Johnson <evan@twiinsen.com>

RUN apt-get update && apt-get install -y git gcc curl make build-essential

COPY src /src
COPY run.sh /bin/run.sh
COPY Makefile.fuzz Makefile.fuzz

#RUN curl http://lcamtuf.coredump.cx/afl/releases/afl-latest.tgz > afl-latest.tgz
COPY afl-latest.tgz afl-latest.tgz


CMD ["/bin/run.sh"] 




