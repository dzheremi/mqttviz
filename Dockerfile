FROM golang:1.13

WORKDIR /go/src/mqttviz
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

EXPOSE 5555

CMD ["mqttviz"]
