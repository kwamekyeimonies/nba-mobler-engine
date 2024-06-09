FROM golang:1.22 AS builder
LABEL maintainer="Mobler"
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . ./
RUN  go build -o /main

# Deploy golang binary
FROM golang:1.22
WORKDIR /
COPY --from=builder /main .
COPY config ./config
COPY logs ./logs
EXPOSE 8976
ENTRYPOINT ["/main"]