from golang:1.19

WORKDIR /usr/src/app

COPY go.mod go.sum ./

RUN apt-get update && apt-get install make && \
    curl -s https://packagecloud.io/install/repositories/golang-migrate/migrate/script.deb.sh | bash && \
    apt-get update && apt-get install migrate && \ 
    go mod download && \
    go mod verify

COPY . .

RUN go build -v -o /usr/local/bin/app ./cmd/main.go

CMD ["app"]