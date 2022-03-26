FROM golang:1.18-bullseye

ADD . /usr/local/go/src/blogtask
WORKDIR /usr/local/go/src/blogtask

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .


COPY *.go .

RUN go build -o /blogapp

ENV pguser="postgres"
ENV pguserpassword="postgres"
ENV pgname="postgres"

EXPOSE 8000
CMD [ "/blogapp" ]