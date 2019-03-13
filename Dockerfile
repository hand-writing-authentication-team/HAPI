FROM golang:1.10

WORKDIR ${GOPATH}/src/github.com/hand-writing-authentication-team/HAPI

ADD . .

RUN go get -u github.com/kardianos/govendor
RUN govendor sync

EXPOSE ${PORT}

CMD ["go", "run", "server.go"]