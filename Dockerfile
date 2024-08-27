#build stage
FROM golang:alpine AS builder
RUN apk add --no-cache git
WORKDIR /go/src/app
COPY . .
RUN go get -d -v ./...
RUN go build -o /go/bin/app ./api/cmd/main.go

#final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /go/bin/app /app
RUN chmod +x /app
CMD ["/app"]
LABEL Name=bootcampapi Version=1.0.0
EXPOSE 8080
