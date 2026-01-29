FROM golang:1.25.5-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o main ./cmd

EXPOSE 8080

CMD [ "./main" ]