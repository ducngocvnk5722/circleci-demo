FROM golang:1.9.1 AS builder

WORKDIR /go/src/github.com/ducngocvnk57/circleci-demo
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o demo-app

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /go/src/github.com/ducngocvnk57/circleci-demo/demo-app .
COPY --from=builder /go/src/github.com/ducngocvnk57/circleci-demo/goose .
COPY --from=builder /go/src/github.com/ducngocvnk57/circleci-demo/run.sh .
RUN chmod +x run.sh
COPY --from=builder /go/src/github.com/ducngocvnk57/circleci-demo/migrations ./migrations/
EXPOSE 8080

ENTRYPOINT ["./run.sh"]