# Build stage
FROM golang:1.24-alpine AS builder
LABEL authors="sainp7"
WORKDIR /builder

# Create ARGs
ARG upx_version=4.2.4
ARG TARGETARCH

# Install required packages, download and install UPX in a single layer
RUN apk add --no-cache xz bash && \
    wget -O /tmp/upx.tar.xz https://github.com/upx/upx/releases/download/v${upx_version}/upx-${upx_version}-${TARGETARCH}_linux.tar.xz && \
    tar -xf /tmp/upx.tar.xz -C /tmp && \
    cp /tmp/upx-${upx_version}-${TARGETARCH}_linux/upx /usr/local/bin/ && \
    chmod +x /usr/local/bin/upx && \
    rm -rf /tmp/*

# set enviornment variables
ENV GO111MODULE=on CGO_ENABLED=0

# download dependencies
COPY go.mod go.sum /builder/
RUN go mod download

#build binary and pack it
COPY . .
RUN go build \
    -ldflags "-s -w" \
    -o /builder/main/out /builder/cmd/main.go && \
    upx -9 /builder/main/out

# Runner stage
FROM alpine:3.20
LABEL authors="sainp7"
RUN apk add --no-cache tzdata netcat-openbsd
WORKDIR /app
COPY --from=builder /builder/main/out main
CMD ["./main"]
