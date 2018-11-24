FROM golang

COPY . .

RUN go build hello-world.go

