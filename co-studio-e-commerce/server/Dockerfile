FROM golang:1.21-alpine as builder

WORKDIR /usr/src/app

ENV DB_HOST=localhost
ENV DB_NAME=postgres
ENV DB_PASSWORD=mysecretpassword
ENV DB_PORT=5432
ENV DB_USER=postgres
ENV PORT=8000
ENV sslmode=disable

COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN GOENV=production CGO_ENABLED=0 GOOS=linux go build -o co-studio-e-commerce

FROM alpine:3.18.3 as production

WORKDIR /usr/src/app

COPY --from=builder /usr/src/app/co-studio .

CMD ["./co-studio-e-commerce"]