FROM golang:1.26.6-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main ./cmd

EXPOSE 8081

CMD [ "./main" ]