FROM golang:1.17.1-alpine3.13

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY ./ ./

RUN go build -o ./pegawai-backend

EXPOSE 1323

CMD ["./pegawai-backend"]