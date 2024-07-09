FROM golang:1.22.5 as base

RUN apt-get update && apt-get install -y ca-certificates openssl

ARG cert_location=/usr/local/share/ca-certificates

# Get certificate from "github.com"
RUN openssl s_client -showcerts -connect github.com:443 </dev/null 2>/dev/null|openssl x509 -outform PEM > ${cert_location}/github.crt
# Get certificate from "proxy.golang.org"
RUN openssl s_client -showcerts -connect proxy.golang.org:443 </dev/null 2>/dev/null|openssl x509 -outform PEM >  ${cert_location}/proxy.golang.crt
# Update certificates
RUN update-ca-certificates

WORKDIR /app

COPY . .

RUN go get -v -t -d ./...

RUN CGO_ENABLED=0 GOOS=linux go build -o main 

FROM scratch as app
COPY --from=base /app .
EXPOSE 8888
ENTRYPOINT [ "/main" ]