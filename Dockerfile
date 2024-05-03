FROM golang:1.22.2-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# CMD ["go", "run", "cmd/main.go"]
CMD ["air"]