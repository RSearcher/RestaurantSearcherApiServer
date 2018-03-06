FROM golang:1.9

ENV APPNAME RestaurantSearcherAPI

RUN mkdir -p /go/src/$APPNAME
WORKDIR /go/src/$APPNAME

ADD . /go/src/$APPNAME

RUN go get -v
