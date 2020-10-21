package main

import (
	"log"

	"github.com/go-pg/migrations/v8"
)

func init() {
	migrations.MustRegisterTx(
		func(db migrations.DB) error {
			log.Println("creating table role_model.permissions...")
			_, err := db.Exec(
				`
				CREATE TABLE role_model.permissions (
					id serial NOT NULL,
					perm_name varchar NOT NULL,
					CONSTRAINT permissions_pk PRIMARY KEY (id),
					CONSTRAINT permissions_un UNIQUE (perm_name)
				);				
			`)
			return err
		}, func(db migrations.DB) error {
			log.Println("dropping table role_model.permissions...")
			_, err := db.Exec(`DROP table role_model.permissions`)
			return err
		})
}
