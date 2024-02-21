FROM golang:1.22-alpine as build 

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main cmd/api/main.go

FROM alpine:3.19.0 as runtime
WORKDIR /app
COPY --from=build /app/main /app/main
COPY --from=build /app/.env /app/.env
COPY --from=build /app/cmd/web /app/cmd/web
COPY cmd/web cmd/web
EXPOSE ${PORT}
CMD ["./main"]