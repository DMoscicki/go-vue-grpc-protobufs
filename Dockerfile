FROM golang:1.20.2-alpine as build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy

COPY .. .

RUN go build -o /backend

FROM alpine

WORKDIR /app

COPY --from=build ./backend /app/backend

EXPOSE 8080/tcp

CMD [ "/app/backend" ]