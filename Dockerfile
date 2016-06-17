FROM golang:latest

MAINTAINER ikether ikether@126.com

RUN mkdir -p /go/src/github.com/zjwdmlmx/ssensor
RUN mkdir -p /var/lib/ssensor
RUN apt-get update
RUN apt-get -y install sqlite
COPY ./ /go/src/github.com/zjwdmlmx/ssensor
RUN go get -d -v github.com/zjwdmlmx/ssensor
RUN go install -v github.com/zjwdmlmx/ssensor
RUN sqlite3 /var/lib/ssensor/data.db < /go/src/github.com/zjwdmlmx/ssensor/support/db.sql

EXPOSE 5050

CMD ["ssensor"]
