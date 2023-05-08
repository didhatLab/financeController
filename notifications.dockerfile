FROM golang:1.19

WORKDIR /app

COPY /go.mod ./
RUN go mod download

COPY . ./

EXPOSE 4005

CMD [ "go", "run", "notifications/main.go" ]
