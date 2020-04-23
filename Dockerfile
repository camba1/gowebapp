FROM golang

WORKDIR /go/src/goWebApp

RUN go get github.com/githubnemo/CompileDaemon

ENV GOPATH="/gp/src/goWebAapp" GOBIN="/gp/src/goWebAapp/bin"

COPY . .

RUN go get -d  -v ./...
RUN go build -o goWebAppLin


CMD ["./goWebAppLin"]
#ENTRYPOINT CompileDaemon --build="go build -o goWebAppLin" --command=./goWebAppLin