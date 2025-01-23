FROM golang:1.23-alpine as builder

WORKDIR /app

COPY go.mod ./

COPY *.go ./

COPY apis-v2 ./apis-v2/

RUN go build -o devops-project .


# run devops-project
FROM alpine:latest 

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/devops-project .

EXPOSE 3456

CMD ["./devops-project"]