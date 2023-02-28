from golang:1.19-alpine AS build

WORKDIR /usr/src/app

COPY go.mod . 
COPY go.sum .

RUN go mod download && \
    go mod verify

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o bin/app cmd/main.go

CMD ["./bin/app"]

FROM golang:1.19-alpine

COPY --from=build /usr/src/app/bin/app .

RUN apk add --no-cache ca-certificates openssl

CMD ["./app"]