FROM golang:1.22-alpine:latest
WORKDIR /bytecrypt_api
RUN go add https://github.com/jackc/pgx
COPY . /bytecrypt_api/
CMD ["go", "run", "."]