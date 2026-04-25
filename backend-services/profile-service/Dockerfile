FROM golang:alpine3.23

RUN apk add --no-cache gcc musl-dev librdkafka-dev pkgconfig

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -tags dynamic -o main ./cmd

EXPOSE 8084

CMD [ "./main" ]