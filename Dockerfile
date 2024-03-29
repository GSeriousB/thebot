# Specify the base image for the go app.
FROM golang:1.17
# Specify that we now need to execute any commands in this directory.
WORKDIR $GOPATH/src/github.com/hakkei
# Add Maintainer Info
LABEL maintainer="tradebot LLC <shubham@tradebot.com>"
# Copy everything from this project into the filesystem of the container.
# COPY . .
# # Obtain the package needed to run code. Alternatively use GO Modules.
# COPY go.mod go.sum ./
# # Compile the binary exe for our app.
# RUN go build -o tradebot cmd/admin/main.go
# # Start the application.
# CMD ["./hakkei"]