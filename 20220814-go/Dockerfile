FROM golang

VOLUME /go
RUN apt update && apt install -y vim

RUN go install golang.org/x/tools/cmd/goimports@latest 
RUN go install github.com/rakyll/hey@latest 
RUN go install golang.org/x/lint/golint@latest

