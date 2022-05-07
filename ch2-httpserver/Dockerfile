FROM golang:1.18.1 as build
MAINTAINER Tanimodori

WORKDIR /home/ 
COPY ./go.mod ./main.go ./
RUN go build -o ./build/httpserver main.go 

FROM debian:stable-slim 
WORKDIR /home/ 
COPY --from=build /home/build/httpserver /home/
RUN chmod 755 ./httpserver
ENTRYPOINT ./httpserver 
EXPOSE 8888
