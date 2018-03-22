FROM        golang:1.8
MAINTAINER  Richard Pape <r1ch4rdp4p3@gmail.com>

ENV     PORT  8080

# Setting up working directory
WORKDIR     /go/src/gin-container
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
ENTRYPOINT  ["/usr/local/bin/go"]
CMD     ["run", "main.go"]


# run with
# "docker build -t microservice-task-app ."
# inside the app applications root directory
