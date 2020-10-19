package datastore

import (
	"github.com/go-pg/pg/v10"
)

// DataBase База данных. Обёртка нужна для того, чтобы создавать методы для подключения к БД
type DataBase struct {
	*pg.DB
}
