FROM golang:1.3

RUN go get github.com/go-martini/martini && go get github.com/martini-contrib/render && go get gopkg.in/mgo.v2 && go get github.com/martini-contrib/binding

ADD . /go/src/github.com/kpurdon/go-todo
WORKDIR /go/src/github.com/kpurdon/go-todo
