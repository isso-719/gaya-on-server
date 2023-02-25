FROM golang:1.20-alpine

WORKDIR /go/src/app
COPY . .

RUN apk upgrade --update && apk --no-cache add git

RUN go get -u github.com/cosmtrek/air && go build -o /go/bin/air github.com/cosmtrek/air

EXPOSE 8080

CMD ["air", "-c", ".air.toml"]