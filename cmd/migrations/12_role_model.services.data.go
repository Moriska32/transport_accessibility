package main

import (
	"log"

	"github.com/go-pg/migrations/v8"
)

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		log.Println("Inserting data into role_model.services")
		_, err := db.Exec(`
		INSERT INTO role_model.services (id, "path", main_path, api_version_path)
			VALUES(1, '/*', '/api', '');
		`)
		return err
	}, func(db migrations.DB) error {
		log.Println("deleting all rows from role_model.services...")
		_, err := db.Exec(`delete from role_model.services`)
		return err
	})
}
