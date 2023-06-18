FROM golang:1.20-buster

RUN apt update && apt install

WORKDIR /go/src/app
COPY ./gocook .
RUN go mod download
RUN go install

CMD ["gocook"]
EXPOSE 8000