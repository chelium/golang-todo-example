FROM golang

RUN mkdir -p /go/src/github.com/chelium/golang-todo-example

ADD . /go/src/github.com/chelium/golang-todo-example

RUN curl https://glide.sh/get | sh
RUN go get  github.com/canthefason/go-watcher
RUN go install github.com/canthefason/go-watcher/cmd/watcher

RUN cd /go/src/github.com/chelium/golang-todo-example && glide install

ENTRYPOINT  watcher -run github.com/chelium/golang-todo-example/todo/cmd -watch github.com/chelium/golang-todo-example/todo
