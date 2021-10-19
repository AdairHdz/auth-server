FROM golang:1.16.4-alpine

WORKDIR /app

COPY . .

RUN go clean --modcache

RUN go build -o main

CMD [ "/app/main" ]
