FROM golang:1.15.5-alpine

RUN apk update && apk add make

WORKDIR ./tasker
COPY . .

RUN make build
