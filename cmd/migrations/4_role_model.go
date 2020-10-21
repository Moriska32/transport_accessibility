package main

import (
	"log"

	"github.com/go-pg/migrations/v8"
)

func init() {
	migrations.MustRegisterTx(
		func(db migrations.DB) error {
			log.Println("creating schema role_model...")
			_, err := db.Exec(`CREATE schema role_model`)
			return err
		}, func(db migrations.DB) error {
			log.Println("dropping schema role_model...")
			_, err := db.Exec(`DROP SCHEMA role_model CASCADE`)
			return err
		},
	)
}
