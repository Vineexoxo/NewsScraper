FROM golang:1.24.2 as builder
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
COPY backend/main.go ./


RUN go mod download
RUN go build -o crawler .
FROM debian:bookworm
WORKDIR /app


RUN apt-get update && apt-get install -y \
    wget gnupg unzip fonts-liberation libnss3 libxss1 libasound2 libatk1.0-0 libatk-bridge2.0-0 libgtk-3-0 \
    && rm -rf /var/lib/apt/lists/*

# Install headless Chromium
RUN wget https://dl.google.com/linux/direct/google-chrome-stable_current_amd64.deb && \
    apt-get update && \
    apt-get install -y ./google-chrome-stable_current_amd64.deb && \
    rm google-chrome-stable_current_amd64.deb

ENV ROD_BROWSER_BIN=/usr/bin/google-chrome
COPY --from=builder /app/crawler /crawler
ENTRYPOINT ["/crawler"]
