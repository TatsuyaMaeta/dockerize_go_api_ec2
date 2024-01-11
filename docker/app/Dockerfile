FROM golang:latest
RUN mkdir go_docker_ec2
WORKDIR /go_docker_ec2


# docker composeのvolumesと揃える
# RUN mkdir go_docker_ec2
# RUN mkdir go
# RUN mkdir go/src
# RUN mkdir go/src/go_docker_ec2 
# WORKDIR /go/src/go_docker_ec2

RUN apt-get update \
    && apt-get install -y mariadb-client

COPY . .

RUN CGO_ENABLED=0 go build -o main main.go

CMD ["./main"]

