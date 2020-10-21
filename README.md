# Сборка серверного приложения
```
export CGO_ENABLED=0 && go build -ldflags "-s -w" -o server_app -gcflags "all=-trimpath=$GOPATH" main.go
```

# Сборка приложения миграции схемы данных и статичных данных базы данных
```
cd cmd/migrations && export CGO_ENABLED=0 && go build -ldflags "-s -w" -o migrations_app -gcflags "all=-trimpath=$GOPATH" && cd -
```