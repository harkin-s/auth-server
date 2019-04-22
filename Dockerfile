FROM golang:alpine
RUN mkdir /app 
ADD ./src /app/
WORKDIR /app
RUN apk update && apk add --no-cache git
RUN go get -d -v ./...
RUN go build -o main .
RUN adduser -S -D -H -h /app appuser
USER appuser
EXPOSE 8080
CMD ["./main"]