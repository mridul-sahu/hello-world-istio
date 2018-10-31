FROM gcr.io/crispy-206704/base:1.1

RUN go get -v bitbucket.org/crispyapp/services/deployment/hello

WORKDIR /go/src/bitbucket.org/crispyapp/services/deployment/hello

COPY . .

RUN ./build.sh

EXPOSE 3000

CMD [ "hello" ]
