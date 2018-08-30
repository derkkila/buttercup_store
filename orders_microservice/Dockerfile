FROM golang:latest
RUN mkdir /app
ADD . /app/
WORKDIR /app

EXPOSE 4201

RUN go build -o main .
CMD ["/app/main"]
