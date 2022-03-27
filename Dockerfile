FROM golang:1.18-bullseye as build

ADD . /usr/local/go/src/blogtask
WORKDIR /usr/local/go/src/blogtask

COPY . .
RUN go mod download


RUN CGO_ENABLED=0 GOOS=linux go build -o blogapp

FROM scratch

COPY --from=build /usr/local/go/src/blogtask/blogapp .

ENV pguser="postgres"
ENV pguserpassword="postgres"
ENV pgname="postgres"

EXPOSE 80

CMD [ "/blogapp" ]

