FROM golang:1.13

ENV GOPATH=/go
ENV GO111MODULE=on
ENV APP=/go/src/othala
ENV PROJECT_ROOT=$APP

WORKDIR $APP
COPY . $APP

RUN go get github.com/githubnemo/CompileDaemon
entrypoint CompileDaemon -log-prefix=false -build="go build -o othala -v app/main.go" -command="./othala"
