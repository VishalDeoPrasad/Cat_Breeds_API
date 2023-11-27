FROM golang:1.21.4-alpine3.18 AS builder
 
#copy entire code into the container, crate a directory called app
WORKDIR /app
 
#whatever command we do here all perform on this container only
# download all dependency
COPY go.mod . 
COPY go.sum . 
 
#donwload all dependency from go mod file into the container
RUN go mod download

#copy all the code inside the directory
COPY . .  
 
#here we have an option to run the main file but instand we are going to build our file
#we build this file and the name of the build is server
RUN go build -o server cmd/main.go
 
 
FROM scratch

#COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

WORKDIR /app
 
COPY --from=builder /app/server .
COPY --from=builder /app/private.pem .
COPY --from=builder /app/pubkey.pem .
 
CMD [ "./server" ]

#command to build docker image
#here 't' is tag, when we build a image we give a tag' and .(dot) is current directry

#docker build -t job_portal_api:latest .