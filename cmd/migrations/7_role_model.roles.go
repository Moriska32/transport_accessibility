package main

import (
	"log"

	"github.com/go-pg/migrations/v8"
)

func init() {
	migrations.MustRegisterTx(
		func(db migrations.DB) error {
			log.Println("creating table role_model.roles...")
			_, err := db.Exec(
				`
				CREATE TABLE role_model.roles (
					id serial NOT NULL,
					role_name varchar NOT NULL,
					CONSTRAINT roles_pk PRIMARY KEY (id),
					CONSTRAINT roles_un UNIQUE (role_name)
				);
			`)
			return err
		}, func(db migrations.DB) error {
			log.Println("dropping table role_model.roles...")
			_, err := db.Exec(`DROP table role_model.roles;`)
			return err
		})
}
