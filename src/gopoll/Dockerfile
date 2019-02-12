FROM golang:1.10

WORKDIR /go/src/github.com/wyrdnixx/votegovue

COPY . .

RUN ["go", "get", "github.com/githubnemo/CompileDaemon"]
RUN ["go", "get", "github.com/tkanos/gonfig"]
#RUN ["go", "get", "github.com/lib/pq"]

RUN ["go", "get", "github.com/labstack/echo"]
RUN ["go", "get", "github.com/labstack/echo/middleware"]
RUN ["go", "get", "github.com/mattn/go-sqlite3"]


ENTRYPOINT CompileDaemon -log-prefix=false -build="go build ./src/gopoll/" -command="./gopoll"
