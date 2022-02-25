FROM golang:1.17
WORKDIR /usr/src/app
COPY go.mod go.sum ./
RUN go mod download && go mod verify
ENV DB_USER mongoadmin
ENV DB_PASS So1pass1S_2022
ENV DB_HOST db
ENV DB_PORT 27017
COPY . .
RUN go build -v -o /usr/src/app ./...
EXPOSE 9090
CMD ["/usr/src/app/sopes-practica1"