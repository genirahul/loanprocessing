FROM amd64/golang:1.12.7

ENV GIN_MODE=release
ENV GO111MODULE=on
EXPOSE 80

RUN mkdir -p /go/src/loanprocessing
WORKDIR /go/src/loanprocessing
COPY . .

# RUN go mod download
RUN go get -d -v ./...
RUN go install -v ./...
RUN go build webserver.go

EXPOSE 8080
CMD ["bash", "start.sh"]