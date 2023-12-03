FROM golang:alpine

WORKDIR /var/www/html/qdjr-api

ADD go.mod .
ADD go.sum .

RUN go mod download
ADD . .

RUN  go install -mod=mod github.com/githubnemo/CompileDaemon

EXPOSE 8000

ENTRYPOINT CompileDaemon --build="go build main.go" --command=./main