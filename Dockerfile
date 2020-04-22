FROM golang

WORKDIR /go/src/goWebApp

ENV GOPATH="/gp/src/goWebAapp" GOBIN="/gp/src/goWebAapp/bin"

COPY . .

RUN go get -d  -v ./...
RUN go build


CMD ["./goWebApp"]
