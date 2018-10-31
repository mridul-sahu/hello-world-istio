FROM golang:1.9.1-alpine

RUN apk add --no-cache git

RUN go get -v github.com/mridul-sahu/hello-world-istio

WORKDIR /go/src/github.com/mridul-sahu/hello-world-istio

COPY . .

RUN go get -v

EXPOSE 3000

CMD [ "hello-world-istio" ]
