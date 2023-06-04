FROM golang:1.20-alpine

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod tidy
RUN go mod download

COPY . ./

RUN go build -o ./gin-api

EXPOSE 5000

CMD [ "./gin-api" ]