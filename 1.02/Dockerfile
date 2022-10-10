FROM golang:1.16-alpine

WORKDIR /app
COPY go.mod ./
COPY go.sum ./
COPY .env ./
RUN go mod download
COPY *.go ./

RUN go build -o /start-server

EXPOSE 8080

CMD [ "/start-server" ]