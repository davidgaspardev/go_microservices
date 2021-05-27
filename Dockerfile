FROM golang:1.16.3-stretch

ENV DIR_APP=/app
ENV GO111MODULE=off

WORKDIR $DIR_APP

COPY . . 

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags '-w' -o microservice

EXPOSE 8080

CMD [ "/app/microservice" ]
