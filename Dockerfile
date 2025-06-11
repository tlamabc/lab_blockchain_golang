FROM golang:1.24.4-bookworm

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o node ./cmd/node

EXPOSE 2201 2202 2203

CMD ["./node"]
