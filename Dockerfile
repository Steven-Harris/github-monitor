FROM golang:1.22.5 as base

WORKDIR /app

COPY . .

RUN go get -v -t -d ./...

RUN CGO_ENABLED=0 GOOS=linux go build -o main 

FROM scratch as app
COPY --from=base /app .
EXPOSE 8888
ENTRYPOINT [ "/main" ]