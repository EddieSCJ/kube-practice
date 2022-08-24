FROM --platform=linux/amd64 golang:alpine

WORKDIR /go/src/app
ADD . .

RUN go mod init
RUN go build -o server .
CMD ["./server"]