FROM golang:1.9-alpine3.7

WORKDIR /go/src/app
RUN apk update && apk add git gcc

COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["go-wrapper", "run"] # ["app"]
