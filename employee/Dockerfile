# golang image where workspace (GOPATH) configured at /go.
FROM golang:1.6

# Install dependencies
RUN go get gopkg.in/mgo.v2
RUN go get github.com/gorilla/mux

# copy the local package files to the container workspace
ADD . /go/src/hr-management/employee

# Setting up working directory
WORKDIR /go/src/hr-management/employee

# Build the users command inside the container.
RUN go install hr-management/employee

# Run the users microservice when the container starts.
ENTRYPOINT /go/bin/users

# Service listens on port 8080.
EXPOSE 8080
