package configuration

import (
	"crypto/tls"

	"github.com/go-pg/pg/v10"
)

// PGConf Конфигурация подключения к базе данных
type PGConf struct {
	Host      string      `json:"host"`
	Port      string      `json:"port"`
	Database  string      `json:"database"`
	Name      string      `json:"name"`
	User      string      `json:"user"`
	Password  string      `json:"password"`
	EnableTLS *tls.Config `json:"enable_tls"`
}

// Connect Подключение к базе данных
func (pgc *PGConf) Connect() *pg.DB {
	return pg.Connect(&pg.Options{
		Addr:      pgc.Host + ":" + pgc.Port,
		User:      pgc.User,
		Password:  pgc.Password,
		Database:  pgc.Database,
		TLSConfig: pgc.EnableTLS,
	})
}
