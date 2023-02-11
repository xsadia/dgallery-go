from golang:1.19 AS build

WORKDIR /usr/src/app

COPY go.mod . 
COPY go.sum .

RUN apt-get update && apt-get install make && \
    curl -s https://packagecloud.io/install/repositories/golang-migrate/migrate/script.deb.sh | bash && \
    apt-get update && apt-get install migrate && \ 
    go mod download && \
    go mod verify

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o bin/app cmd/main.go

FROM golang:1.19-alpine

COPY --from=build /usr/src/app/bin/app .

CMD ["./app"]