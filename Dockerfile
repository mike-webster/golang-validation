FROM golang:1.11-alpine3.8

RUN apk add --no-cache ca-certificates git make curl mysql-client gcc musl-dev

WORKDIR /golang-validation

ENV GO111MODULE=on

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN ["go", "get", "github.com/githubnemo/CompileDaemon"]
#RUN GOOS=linux go build -o golang-validation .

# For Web
EXPOSE 3001

ENTRYPOINT CompileDaemon -build="go build -o golang-validator ." -command="./golang-validator"

#ENTRYPOINT ["./golang-validation"]