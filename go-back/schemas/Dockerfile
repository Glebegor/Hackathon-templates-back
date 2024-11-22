# Base image
FROM ubuntu:20.04

# Metadata
LABEL author="Glebegor"

# Set environment variables
ENV DEBIAN_FRONTEND=noninteractive
ENV LANG=C.UTF-8

# Install required packages
RUN apt-get update && apt-get install -y \
    make \
    curl \
    iputils-ping \
    postgresql-client \
    && rm -rf /var/lib/apt/lists/*

# Install golang-migrate
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.tar.gz | tar xvz && \
    mv migrate /usr/local/bin/

# Copy project files
WORKDIR /app
COPY . .

# Set default command
CMD ["make", "migrate-up"]