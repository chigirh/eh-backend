FROM golang:1.18.1

ENV GO111MODULE=on
ENV ROOT=/go/src/eh-backend-api

WORKDIR ${ROOT}

COPY go.mod ./
COPY go.sum ./
COPY ./ ./

RUN go mod download

EXPOSE 8080

# RUN go run eh-backend-application.go

RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

CMD ["air"]