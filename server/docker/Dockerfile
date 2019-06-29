FROM golang

MAINTAINER makemoney

RUN go get -u "github.com/astaxie/beego" \
    && go get -u "github.com/bitly/go-simplejson" \
    && go get -u "github.com/go-sql-driver/mysql" \
    && go get -u "github.com/make-money-sysu/Make-Money"

RUN cd /go/src/github.com/make-money-sysu/Make-Money/front-end/make-money \
    && npm install \
    && npm run build \
    && mv dist/* ../../server/static/ 

WORKDIR /go/src/github.com/make-money-sysu/Make-Money/server

CMD ["go", "run", "main.go"]