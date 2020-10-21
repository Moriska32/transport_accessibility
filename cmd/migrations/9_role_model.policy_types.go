package main

import (
	"log"

	"github.com/go-pg/migrations/v8"
)

func init() {
	migrations.MustRegisterTx(
		func(db migrations.DB) error {
			log.Println("creating table role_model.policy_types...")
			_, err := db.Exec(
				`
				CREATE TABLE role_model.policy_types (
					id serial NOT NULL,
					policy_name varchar NOT NULL,
					CONSTRAINT policy_types_pk PRIMARY KEY (id),
					CONSTRAINT policy_types_un UNIQUE (policy_name)
				);
			`)
			return err
		}, func(db migrations.DB) error {
			log.Println("dropping table role_model.policy_types...")
			_, err := db.Exec(`DROP table role_model.policy_types`)
			return err
		})
}
