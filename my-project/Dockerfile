FROM golang:1.11.2 as dev
MAINTAINER shubham jagdhane

# Install Dependecies
RUN go get github.com/go-sql-driver/mysql
RUN go get github.com/golang-migrate/migrate
RUN go get github.com/golang-migrate/migrate/database/mysql
RUN go get github.com/golang-migrate/migrate/source/file
RUN go get github.com/gorilla/mux

ADD . /go/src/github.com/shubham1010/mysql-nginx-docker/my-project

WORKDIR /go/src/github.com/shubham1010/mysql-nginx-docker/my-project

#build binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:3.9.2 as prod
MAINTAINER shubham jagdhane

WORKDIR /root/

# copy main binary
COPY --from=dev /go/src/github.com/shubham1010/mysql-nginx-docker/my-project .

EXPOSE 8080
ENTRYPOINT ["./main"]
