package main

import (
	"log"

	"github.com/go-pg/migrations/v8"
)

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		log.Println("Inserting data into role_model.policy_types")
		_, err := db.Exec(`
		INSERT INTO role_model.policy_types (id, policy_name)
		VALUES(1, 'p');
		INSERT INTO role_model.policy_types (id, policy_name)
		VALUES(2, 'g');
		`)
		return err
	}, func(db migrations.DB) error {
		log.Println("deleting all rows from role_model.policy_types...")
		_, err := db.Exec(`delete from role_model.policy_types`)
		return err
	})
}
