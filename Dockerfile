# syntax=docker/dockerfile:1

FROM golang:1.18.4-alpine

WORKDIR /src

COPY . .
RUN go mod tidy
RUN go mod download

COPY . ./

RUN go build -o main .

EXPOSE 8080

CMD [ "./main" ]