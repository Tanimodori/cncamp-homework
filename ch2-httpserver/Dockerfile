FROM golang:1.18.1 as build
MAINTAINER Tanimodori

ENV GO111MODULE=on
WORKDIR /home/ 
ADD . /home/ 
RUN go get -u -v
RUN go build -o ./build/httpserver main.go 

FROM debian:stable-slim 
WORKDIR /home/ 
COPY --from=build /home/build/httpserver /home/
RUN chmod 755 ./httpserver
ENTRYPOINT ./httpserver 
EXPOSE 8080
