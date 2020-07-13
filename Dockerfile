FROM golang:1.13

ENV GOPATH=/go
ENV GO111MODULE=on
ENV APP=/go/src/othala
ENV PROJECT_ROOT=$APP

WORKDIR $APP
COPY . $APP

RUN pwd
RUN ls
RUN go get github.com/githubnemo/CompileDaemon
