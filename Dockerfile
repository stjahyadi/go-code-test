FROM golang:1.14

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN go get -u github.com/gorilla/mux && go get -u github.com/gorilla/handlers && go get -u github.com/prometheus/client_golang/prometheus/promhttp && go get -u github.com/dgrijalva/jwt-go

RUN go build -o main .

EXPOSE 8081

CMD ["/app/main"]
#CMD ["sh", "-c", "tail -f /dev/null"]