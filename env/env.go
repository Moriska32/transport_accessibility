package env

import (
	"transport_accessibility_bus/configuration"

	"github.com/go-pg/pg/v10"
)

// PGInstances Набор подключений к Postgres
type PGInstances map[string]*pg.DB

// Env Для Dependency injection (внедрение зависимостей)
type Env struct {
	DBs PGInstances
}

// PrepareEnv Устанавливаем соединение со всеми экземплярами баз данных Postgres и Redis
func PrepareEnv(cfg *configuration.AppConfiguration) *Env {
	env := &Env{
		DBs: make(PGInstances),
	}
	for i := range cfg.PostgresConfiguration {
		env.DBs[cfg.PostgresConfiguration[i].Name] = cfg.PostgresConfiguration[i].Connect()
	}
	return env
}
