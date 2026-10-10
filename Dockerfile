# build stage
FROM golang:1.27.1-alpine3.24 AS builder

WORKDIR /app 

COPY . .

RUN go build -o main main.go

RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.18.3


# run stage

FROM alpine:3.24

WORKDIR /app

COPY --from=builder /app/main .

COPY --from=builder /go/bin/migrate ./migrate

COPY app.env .
COPY --chmod=755 start.sh .

COPY db/migration ./migration


EXPOSE 8080

CMD ["/app/main"]

ENTRYPOINT ["/app/start.sh"]

