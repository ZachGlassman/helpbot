FROM golang:1.11.1

WORKDIR /go/src/app
COPY . .
RUN go get -d -v ./...
RUN go install -v ./...

cmd ["app"]
