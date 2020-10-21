package main

import (
	"log"

	"github.com/go-pg/migrations/v8"
)

func init() {
	migrations.MustRegisterTx(
		func(db migrations.DB) error {
			log.Println("creating extention postgis...")
			_, err := db.Exec(`CREATE EXTENSION IF NOT EXISTS postgis;`)
			return err
		}, func(db migrations.DB) error {
			log.Println("dropping extention postgis...")
			_, err := db.Exec(`DROP EXTENSION postgis;`)
			return err
		},
	)
}
