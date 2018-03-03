FROM golang:1.8

ENV APPNAME RestaurantSearcherAPI

ADD . /go/src/$APPNAME
WORKDIR /go/src/$APPNAME

EXPOSE 8080

RUN go get

CMD ["go", "run", "main.go"]
