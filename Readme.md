# Calculator GRPC Service

## System Requirements :

### 1. Go must be installed in the system. To install go please follow

    https://golang.org/doc/install

## To Run the service, please follow below steps:

### 1. Go to server folder and run the 'server.go' file using

    go run server.go

### 2. Go to client folder and run

    go build .

### This will create 'client' executable, then you can run

    ./client -method add -a 2 -b 4

### And you will see output in format

    Response from server: 6

### Available methods are : add, sub, div, mul
