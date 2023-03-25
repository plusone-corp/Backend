# Build Stage
# First pull Golang image
FROM golang:latest as build-env

# Set environment variable
ENV APP_NAME Backend
ENV ENV_NAME .env
ENV CMD_PATH main.go

# Copy application data into image
COPY ./ $GOPATH/src/$APP_NAME
COPY ./.env /$ENV_NAME
WORKDIR $GOPATH/src/$APP_NAME

# Budild application
RUN CGO_ENABLED=0 go build -v -o /$APP_NAME $GOPATH/src/$APP_NAME/$CMD_PATH

# Run Stage
FROM alpine:latest

# Set environment variable
ENV APP_NAME Backend
ENV ENV_NAME .env

# Copy only required data into this image
COPY --from=build-env /$APP_NAME .
COPY --from=build-env /$ENV_NAME .

EXPOSE 80

# Start app
CMD ./$APP_NAME