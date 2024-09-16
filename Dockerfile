#build stage
FROM golang:1.16-alpine3.13 AS builder
WORKDIR /go/src/app
COPY . .
RUN go build -o /go/bin/app ./api/cmd/main.go


#final stage
FROM alpine3.13
WORKDIR app/
COPY --from=builder /app/main.go .
EXPOSE  8080

CMD ["app/main"]

