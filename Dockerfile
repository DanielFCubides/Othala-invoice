FROM golang:1.18 AS builder

ENV GOPATH=/go
ENV GO111MODULE=on
ENV APP=/go/src/othala
ENV PROJECT_ROOT=$APP

WORKDIR $APP
COPY . $APP

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o othala app/main.go

FROM alpine:latest
ENV APP=/go/src/othala
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder $APP/othala ./
CMD ["./othala"]