package main

import (
	"log"

	"github.com/go-pg/migrations/v8"
)

func init() {
	migrations.MustRegisterTx(
		func(db migrations.DB) error {
			log.Println("creating extention uuid...")
			_, err := db.Exec(`CREATE extension if not exists "uuid-ossp"`)
			return err
		}, func(db migrations.DB) error {
			log.Println("dropping extention uuid...")
			_, err := db.Exec(`DROP EXTENSION "uuid-ossp"`)
			return err
		},
	)
}
