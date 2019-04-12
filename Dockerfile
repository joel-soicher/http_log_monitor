FROM golang:1.11

run touch /tmp/access.log

WORKDIR $GOPATH/src/github.com/joel-soicher/http_log_monitor
COPY . .

RUN go get -v github.com/hpcloud/tail
RUN go install github.com/joel-soicher/http_log_monitor/src/http_log_monitor

ENTRYPOINT ["http_log_monitor"]
CMD [""]
