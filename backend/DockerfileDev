FROM golang:1.17.1-alpine3.13

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download
RUN go get github.com/pilu/fresh

COPY ./ ./

EXPOSE 1323

CMD ["fresh"]