package main

import (
	"log"

	"github.com/go-pg/migrations/v8"
)

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		log.Println("Inserting data into role_model.permissions")
		_, err := db.Exec(`
		INSERT INTO role_model.permissions (id, perm_name)
			VALUES(1, 'deny');
		INSERT INTO role_model.permissions (id, perm_name)
			VALUES(2, 'allow');
		`)
		return err
	}, func(db migrations.DB) error {
		log.Println("deleting all rows from role_model.permissions...")
		_, err := db.Exec(`delete from role_model.permissions`)
		return err
	})
}
