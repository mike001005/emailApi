# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang

ENV MG_API_KEY=key-233b2d3f9b230f830c8fb9cda8da37c3
ENV MG_DOMAIN=sandboxcdd04810d08f4f088d4bb8f1eef19fd2.mailgun.org
ENV MG_PUBLIC_API_KEY=pubkey-0bf35947560580b92605a957f49351b2
ENV SENDGRID_API_KEY=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

RUN go get -v -u github.com/gorilla/mux
RUN go get -v -u github.com/sendgrid/sendgrid-go
RUN go get -v -u gopkg.in/mailgun/mailgun-go.v1
RUN go get -v -u github.com/goware/emailx

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