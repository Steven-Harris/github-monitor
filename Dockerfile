FROM golang:1.22.5

WORKDIR /app

COPY . .

RUN go get -v -t -d ./...

RUN CGO_ENABLED=0 GOOS=linux go build -o main .

EXPOSE 3000

CMD ["/app/main"]