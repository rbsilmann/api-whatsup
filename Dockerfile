FROM golang:1.19.2-alpine3.16

WORKDIR /app

COPY . ./

RUN go build -o /server

EXPOSE 9098

CMD [ "/server" ]