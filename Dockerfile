FROM gcr.io/crispy-206704/base:1.1

RUN go get -v github.com/mridul-sahu/hello-world-istio

WORKDIR /go/src/github.com/mridul-sahu/hello-world-istio

COPY . .

RUN ./build.sh

EXPOSE 3000

CMD [ "hello" ]
