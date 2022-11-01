# syntax=docker/dockerfile:1

FROM golang:1.18.4-alpine As builder

WORKDIR /go/src/app

COPY . .

RUN go build -o main .

FROM alpine

WORKDIR /app

COPY --from=builder /go/src/app/ /app/

EXPOSE 8080

CMD [ "./main" ]


# Normal Way
# WORKDIR /go/src

# COPY . .

# RUN go build -o main .

# EXPOSE 8080

# CMD [ "./main" ]