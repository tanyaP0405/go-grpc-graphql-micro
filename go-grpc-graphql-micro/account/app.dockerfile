FROM golang:1.13-alpine3.11 AS build

RUN apk --no-cache add gcc g++ make ca-certificates

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o app ./account/cmd/account

FROM alpine:3.11

WORKDIR /usr/bin

COPY --from=build /app/app .

EXPOSE 8080

CMD ["./app"]