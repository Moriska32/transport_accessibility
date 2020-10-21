package main

import (
	"log"

	"github.com/go-pg/migrations/v8"
)

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		log.Println("Inserting data into role_model.policies")
		_, err := db.Exec(`
		INSERT INTO role_model.policies (id, policy_type_id, policy_role_id, user_id, service_id, permission_id)
			VALUES(1, 1, 1, NULL, 1, 2);
		INSERT INTO role_model.policies (id, policy_type_id, policy_role_id, user_id, service_id, permission_id)
			VALUES(2, 1, 2, NULL, 1, 2);
		INSERT INTO role_model.policies (id, policy_type_id, policy_role_id, user_id, service_id, permission_id)
			VALUES(3, 1, 3, NULL, 1, 2);
		`)
		return err
	}, func(db migrations.DB) error {
		log.Println("deleting all rows from role_model.policies...")
		_, err := db.Exec(`delete from role_model.policies`)
		return err
	})
}
