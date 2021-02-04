FROM golang
MAINTAINER suresh kumar vv <https://github.com/suresh-baswa>
RUN go get github.com/suresh-baswa/suresh
WORKDIR /go/src/app
RUN cp /go/src/github.com/suresh-baswa/suresh/RestServiceExample.go /go/src/app
RUN go build /go/src/app/RestServiceExample.go
CMD ["/go/src/app/RestServiceExample"]