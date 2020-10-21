package main

import (
	"log"

	"github.com/go-pg/migrations/v8"
)

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		log.Println("Inserting data into role_model.roles")
		_, err := db.Exec(`
		INSERT INTO role_model.roles (id, role_name)
			VALUES(1, 'admin');
		INSERT INTO role_model.roles (id, role_name)
			VALUES(2, 'user');
		INSERT INTO role_model.roles (id, role_name)
			VALUES(3, 'developer');		
		`)
		return err
	}, func(db migrations.DB) error {
		log.Println("deleting all rows from role_model.roles...")
		_, err := db.Exec(`delete from role_model.roles`)
		return err
	})
}
