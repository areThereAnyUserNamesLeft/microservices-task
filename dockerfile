FROM        golang:1.8
MAINTAINER  Richard Pape <r1ch4rdp4p3@gmail.com>

ENV
ENV     PORT  8080

# Setting up working directory
WORKDIR     /home/go/src/gin-container
ADD         . /go/src/gin-container

RUN     go get github.com/tools/godep
RUN     go get github.com/gin-gonic/gin
RUN     go get gopkg.in/gorp.v1
RUN     go get github.com/go-sql-driver/mysql
RUN     go install github.com/tools/godep
RUN     go install github.com/gin-gonic/gin

# Restore godep dependencies
#RUN godep restore

EXPOSE 8080
RUN /bin/bash -c "mkdir -p /home/bin/go/microservice-task"

WORKDIR /usr/local/bin/go/microservice-task

COPY . .

CMD ["./main"]

# run with
# "docker build -t microservice-task ."
# inside the app applications root directory
