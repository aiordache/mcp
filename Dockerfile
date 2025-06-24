FROM golang:1.24.4-bookworm AS builder

WORKDIR /mcp
COPY . .


RUN go mod download
RUN go build -o ./build/mcp ./cmd/mcp
RUN chmod +x ./build/mcp

EXPOSE 8080
CMD ["/mcp/build/mcp"]