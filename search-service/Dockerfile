FROM golang:1.21-alpine as build

WORKDIR /go/src/app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /search-service ./main.go

FROM gcr.io/distroless/base-debian11

WORKDIR /go/src/app

COPY --from=build /search-service /search-service

EXPOSE 4002

CMD ["/search-service"]