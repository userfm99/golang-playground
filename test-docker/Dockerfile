FROM golang:alpine as Builder 

ENV GO111MODULE=on

WORKDIR /smartfren

COPY . .

RUN go mod download

RUN go build -o main .

WORKDIR /dist

RUN cp /smartfren/main .

EXPOSE 8080

CMD ["/smartfren/main"]

