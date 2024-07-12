FROM golang:1.22.5 as base

RUN apt-get update && apt-get install -y ca-certificates openssl
RUN mkdir -p /certs
RUN openssl s_client -showcerts -connect api.github.com:443 </dev/null 2>/dev/null | openssl x509 -outform PEM > /certs/api.github.crt
RUN update-ca-certificates

WORKDIR /app

COPY . .

RUN go get -v -t -d ./...

RUN CGO_ENABLED=0 GOOS=linux go build -o main 

FROM scratch as app

COPY --from=base /etc/ssl/certs /etc/ssl/certs
COPY --from=base /certs/api.github.crt /etc/ssl/certs/api.github.crt

COPY --from=base /app .
EXPOSE 8888
ENTRYPOINT [ "/main" ]