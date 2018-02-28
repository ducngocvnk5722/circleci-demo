FROM golang:1.9.1

WORKDIR /go/src/github.com/eb-go-sample

COPY app.go	.

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=0 /go/src/github.com/eb-go-sample/app    .

EXPOSE 3000

CMD ["./app"]