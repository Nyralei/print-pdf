FROM golang:1.22.4 AS builder

RUN apt update && \
    apt install -y libgtk-3-dev

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o print-pdf -v ./cmd/print-pdf

FROM scratch

WORKDIR /root/

COPY --from=builder /app/print-pdf .

CMD ["./print-pdf"]