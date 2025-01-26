FROM golang:latest

WORKDIR /app

COPY . . 

RUN go build -o cmb .

EXPOSE 8088

CMD [ "./cmb release" ]