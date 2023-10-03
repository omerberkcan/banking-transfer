FROM golang:1.19-alpine AS builder


WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download


COPY main.go .
COPY ./internal ./internal
COPY ./model ./model
COPY ./dto ./dto
COPY ./helper ./helper


RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -a -ldflags '-s -w -extldflags "-static"' .

FROM alpine:latest

WORKDIR /app

CMD [ "ls" ]

ENTRYPOINT ["./banking-transfer"]

COPY --from=builder /app/banking-transfer /app/banking-transfer