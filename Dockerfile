FROM golang:1.23

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN go build -v -o /bin/app ./cmd/server

RUN mkdir -p /app/data

RUN go run ./cmd/migrate/main.go --config=./config/docker-server.yml

CMD ["/bin/app"]
