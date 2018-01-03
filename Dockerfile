FROM golang:onbuild


RUN go install github.com/golang/example/outyet
EXPOSE 8081