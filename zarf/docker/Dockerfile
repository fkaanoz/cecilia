FROM golang:1.20.11 AS BUILD
ARG RELEASE

WORKDIR /var/app
COPY . .

RUN GOARCH=amd64 GOOS=linux go build -ldflags="-X main.release=$RELEASE" -o bin/api ./cmd/cecilia/main.go


FROM alpine:3.18.4

WORKDIR /var/app

COPY --from=BUILD /var/app/bin/api /var/app/

CMD ["./api"]