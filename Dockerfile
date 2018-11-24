FROM golang

COPY . .

RUN go build

