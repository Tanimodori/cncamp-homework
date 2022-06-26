FROM golang:1.18.1 as build
MAINTAINER Tanimodori

ENV GO111MODULE=on
WORKDIR /home/ 
COPY ./go.mod ./go.sum ./main.go ./metrics/metrics.go ./
RUN go get -u -v -d
RUN go build -o ./build/httpserver main.go 

FROM debian:stable-slim 
WORKDIR /home/ 
COPY --from=build /home/build/httpserver /home/
RUN chmod 755 ./httpserver
ENTRYPOINT ./httpserver 
EXPOSE 8080
