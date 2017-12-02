FROM golang:latest

ENV GOPATH /app

RUN mkdir /app 
COPY ./** /app/
WORKDIR /app

RUN go build -o gos . 
CMD ["/app/gos"]