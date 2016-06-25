FROM golang:latest

MAINTAINER ikether ikether@126.com

RUN mkdir -p /go/src/github.com/zjwdmlmx/ssensor
RUN mkdir -p /var/lib/ssensor
RUN mkdir -p /var/log/supervisord

RUN apt-get update
RUN apt-get -y install sqlite
RUN apt-get -y install supervisor
RUN apt-get -y install redis-server

COPY ./ /go/src/github.com/zjwdmlmx/ssensor
COPY ./supervisord.conf /etc/supervisord.conf

RUN go get -d -v github.com/zjwdmlmx/ssensor
RUN go get -d -v github.com/zjwdmlmx/ssensor/cleaner
RUN go install -v github.com/zjwdmlmx/ssensor
RUN go install -v github.com/zjwdmlmx/ssensor/cleaner

RUN sqlite3 /var/lib/ssensor/data.db < /go/src/github.com/zjwdmlmx/ssensor/support/db.sql

EXPOSE 5050

VOLUME ["/var/lib/ssensor"]

CMD ["supervisord", "-c", "/etc/supervisord.conf"]
