FROM golang:1.3

ADD . /go/src/github.com/kpurdon/go-blog
WORKDIR /go/src/github.com/kpurdon/go-blog

RUN go get github.com/go-martini/martini && go get github.com/martini-contrib/render && go get gopkg.in/mgo.v2 && go get github.com/martini-contrib/binding
