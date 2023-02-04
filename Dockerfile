FROM golang:1.17
WORKDIR /blog
COPY . .
RUN go mod download
RUN go install -mod=mod github.com/githubnemo/CompileDaemon
ENTRYPOINT CompileDaemon --build="go build -o app ./cmd/server" -polling=true --command=./app