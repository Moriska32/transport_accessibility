# Миграция схемы данных
Необходимо осуществить предустановку PostgreSQL версии 12.x

Так же необходимо предустановить геоинформационную надстройку Postgis версии 3.x

Иницализация миграций:
```shell
./migrations_app --addr АДРЕС_БАЗЫ_ДАННЫХ:ПОРТ_БАЗЫ_ДАННЫХ --db НАИМЕНОВАНИЕ_БАЗЫ_ДАННЫХ --login ЛОГИН_ПОЛЬЗОВАТЕЛЯ --password ПАРОЛЬ_ПОЛЬЗОВАТЕЛЯ init
```

Запуск миграций:
```shell
./migrations_app --addr АДРЕС_БАЗЫ_ДАННЫХ:ПОРТ_БАЗЫ_ДАННЫХ --db НАИМЕНОВАНИЕ_БАЗЫ_ДАННЫХ --login ЛОГИН_ПОЛЬЗОВАТЕЛЯ --password ПАРОЛЬ_ПОЛЬЗОВАТЕЛЯ
```

# Использование серверного приложения
Предварительно нужно заполнить файл [configuration.json](configuration.json) нужными атрибутами.
Так же необходимо сгенерировать RS512 ключи для авторизационной модели. Пример генерации таких ключей:
```shell
ssh-keygen -t rsa -b 4096 -m PEM -f jwtRS512.key
openssl rsa -in jwtRS512.key -pubout -outform PEM -out jwtRS512.key.pub
```

Запуск приложения
```shell
./server_app --conf configuration.json
```
, где флаг _conf_ указывает на файл конфигурации [configuration.json](configuration.json)