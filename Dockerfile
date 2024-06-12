FROM golang:alpine
RUN apk update && apk add bash
WORKDIR /app
COPY go.mod ./
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /server
LABEL Name=asciiartweb Version=0.0.1
CMD ["/server"]
