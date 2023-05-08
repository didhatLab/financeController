FROM golang:1.19

WORKDIR /app

COPY /go.mod ./
RUN go mod download

COPY . ./

EXPOSE 4001

CMD [ "go", "run", "auth/main.go" ]
