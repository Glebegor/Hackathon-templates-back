FROM golang:latest
LABEL author="Glebegor"

WORKDIR /app

RUN apt-get update && apt-get install -y make curl

RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.tar.gz | tar xvz && \
    mv migrate /usr/local/bin/

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /main ./cmd/main.go

CMD ["/main"]