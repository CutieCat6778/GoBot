# Build Stage
# First pull Golang image
FROM golang:latest as build-env

# Set environment variable
ENV APP_NAME geobot
ENV CMD_PATH main.go

# Copy application data into image
COPY ./ $GOPATH/src/$APP_NAME
COPY ./.env $GOPATH/src/$APP_NAME
COPY ./.env /$APP_NAME
WORKDIR $GOPATH/src/$APP_NAME

# Budild application
RUN CGO_ENABLED=0 go build -v -o /$APP_NAME $GOPATH/src/$APP_NAME/$CMD_PATH

# Run Stage
FROM alpine:latest

# Set environment variable
ENV APP_NAME geobot

# Copy only required data into this image
COPY --from=build-env /$APP_NAME .

EXPOSE 3000

# Start app
CMD ./$APP_NAME