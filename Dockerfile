FROM golang:alpine

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY . .

EXPOSE 1323

RUN go mod tidy

RUN go build 

CMD [ "go", "run main.go", "--host", "0.0.0.0" ] 
