# golang image where workspace (GOPATH) configured at /go.
FROM golang:1.6

# Install dependencies
RUN go get github.com/gorilla/mux
RUN go get github.com/go-sql-driver/mysql

# copy the local package files to the container workspace
ADD . /go/src/backend/employee

# Setting up working directory
WORKDIR /go/src/backend/employee

# Build the employee command inside the container.
RUN go install backend/employee

# Run the users microservice when the container starts.
ENTRYPOINT /go/bin/users

# Service listens on port 8080.
EXPOSE 8080
