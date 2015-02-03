FROM golang:1.4

COPY . /gostatic/

RUN cd /gostatic && go build .

RUN mkdir /data

VOLUME ["/data"]

EXPOSE 80

ENTRYPOINT ["/gostatic/gostatic"]