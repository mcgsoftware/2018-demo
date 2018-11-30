FROM golang:latest


# Copy the local package files to the container's workspace
ADD . .

# get dependencies
RUN go get github.com/gorilla/mux

# Command to build the microservice app
RUN go install rccldemo.com/service


# Run this command by default when the container starts.
ENTRYPOINT /go/bin/service


# Document that the hellogcp listens on port 8080.
EXPOSE 8082