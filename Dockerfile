## base image
FROM golang:1.12.0-alpine3.9
## application root directory
RUN mkdir /simple_api
## Copy files into application root directory
ADD . /simple_api
## Set working directory
WORKDIR /simple_api
## Download named modules
RUN go mod download
## Compile the binary, run api 
RUN go build -o main .
## Our start command which kicks off
## our newly created binary executable
CMD ["/simple_api/main"]