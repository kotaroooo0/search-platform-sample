FROM golang:1.21

WORKDIR /app

ADD go.mod go.sum main.go ./
RUN go mod download

ADD main.go ./
RUN go build -o main /app/main.go

ENTRYPOINT ["/app/main"]
