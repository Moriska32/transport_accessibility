# Сборка
```
export CGO_ENABLED=0 && go build -ldflags "-s -w" -o server_app -gcflags "all=-trimpath=$GOPATH" main.go
```