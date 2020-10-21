package main

import (
	"log"

	"github.com/go-pg/migrations/v8"
)

func init() {
	migrations.MustRegisterTx(
		func(db migrations.DB) error {
			log.Println("creating extention pgcrypto...")
			_, err := db.Exec(`CREATE EXTENSION IF NOT EXISTS pgcrypto;`)
			return err
		}, func(db migrations.DB) error {
			log.Println("dropping extention pgcrypto...")
			_, err := db.Exec(`DROP EXTENSION pgcrypto;`)
			return err
		},
	)
}
