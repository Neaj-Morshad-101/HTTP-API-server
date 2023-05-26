# base image go lang is used as a os for running the commands and building ....
FROM golang:latest AS builder
## creates a directory name app inside the container
WORKDIR /app
## copy all files from the local to the current app directory
COPY . .
# ENV key=value
# RUN go build

##need to google this line
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o http-api-server .

# EXPOSE 8080


## multi stage building starts here
## the idea is the avoid build wil be buil in another light weight os alpine 
FROM alpine

WORKDIR /app
## defining flags from to copy form the previous os 
RUN apk add curl
COPY --from=builder /app/http-api-server .
#COPY .env .env

ENTRYPOINT ["./http-api-server"]

CMD ["start"]



#docker build -t neajmorshad/http-api-server:0.0.1 .
#docker push neajmorshad/http-api-server:0.0.1
