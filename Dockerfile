FROM golang:alpine

ARG PRODUCT_NAME="app"
RUN mkdir -p /${PRODUCT_NAME}
WORKDIR /${PRODUCT_NAME}

COPY . .
RUN go mod download

RUN go build -o start_app

ENTRYPOINT  ["./start_app"]