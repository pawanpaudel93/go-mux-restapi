FROM golang:1.14-alpine AS builder

# Updates the repository and installs git
RUN apk update && apk upgrade && apk add --no-cache git

# Enable go modules
ENV GO111MODULE=on

# Switches to /tmp/app as the working directory, similar to 'cd'
WORKDIR /tmp/app

# Note here: To avoid downloading dependencies every time we
# build image. Here, we are caching all the dependencies by
# first copying go.mod and go.sum files and downloading them,
# to be used every time we build the image if the dependencies
# are not changed.

# Copy go mod and sum files
COPY go.mod ./
COPY go.sum ./

# Download all dependencies.
RUN go mod download

# Now, copy the source code
COPY . .

# Note here: CGO_ENABLED is disabled for cross system compilation
# It is also a common best practise.

# Build the application.
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/main .

# The project has been successfully built and we will use a
# lightweight alpine image to run the server 
FROM alpine:latest

# Adds CA Certificates to the image
RUN apk add ca-certificates

# Copy the Pre-built binary file
COPY --from=builder /tmp/app/bin/main /app/main
# COPY --from=builder /tmp/app/.env .

# Switches working directory to /app
WORKDIR "/app"

# Exposes the 5000 port from the container
EXPOSE 5000

# RUN chmod +x ./main
# Run executable
CMD ["./main"]
