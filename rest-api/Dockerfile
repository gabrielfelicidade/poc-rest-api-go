FROM golang:1.13.4-alpine3.10

WORKDIR /go/src/poc-rest-api-go/rest-api
COPY . /go/src/poc-rest-api-go/rest-api

RUN apk --no-cache add curl
RUN apk --no-cache add git

RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

RUN dep ensure

RUN go build
RUN go install

CMD ["rest-api"]
