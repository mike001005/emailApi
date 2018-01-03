# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang

ENV MG_API_KEY=YOUR_KEY
ENV MG_DOMAIN=YOUR_KEY
ENV MG_PUBLIC_API_KEY=YOUR_KEY
ENV SENDGRID_API_KEY=YOUR_KEY

RUN go get -v -u github.com/gorilla/mux
RUN go get -v -u github.com/sendgrid/sendgrid-go
RUN go get -v -u gopkg.in/mailgun/mailgun-go.v1

# Copy the local package files to the container's workspace.
ADD . /go/src/github.com/mike001005/emailApi

RUN go install github.com/mike001005/emailApi
RUN mkdir /app

ADD . /app/
WORKDIR /app
RUN go build -o main .

CMD ["/app/main"]

# Document that the service listens on port 8080.
EXPOSE 8081