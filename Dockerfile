FROM golang:1.19.2-alpine

WORKDIR /home/app

COPY . .

RUN apk update && apk upgrade && apk add bash

CMD ["tail", "-f" , "/dev/null"]