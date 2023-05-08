FROM golang:1.19

WORKDIR /app

COPY /go.mod ./
RUN go mod download

COPY . ./

EXPOSE 4000

CMD [ "go", "run", "finances/main.go" ]
