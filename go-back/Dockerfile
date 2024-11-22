FROM golang:latest
LABEL author="Glebegor"

WORKDIR /app

RUN apt-get update && apt-get install -y make curl iputils-ping postgresql-client && rm -rf /var/lib/apt/lists/*

RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.tar.gz | tar xvz && \
    mv migrate /usr/local/bin/

RUN curl -Lo yq https://github.com/mikefarah/yq/releases/download/v4.34.2/yq_linux_amd64 && \
    chmod +x yq && \
    mv yq /usr/local/bin/


COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /main ./cmd/main.go

CMD ["make", "inside-cluster-run"]