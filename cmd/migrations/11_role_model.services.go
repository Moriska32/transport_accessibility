package main

import (
	"log"

	"github.com/go-pg/migrations/v8"
)

func init() {
	migrations.MustRegisterTx(
		func(db migrations.DB) error {
			log.Println("creating table role_model.services...")
			_, err := db.Exec(
				`
				CREATE TABLE role_model.services (
					id serial NOT NULL,
					"path" varchar NOT NULL,
					main_path varchar NOT NULL DEFAULT '/api'::character varying,
					api_version_path varchar NOT NULL DEFAULT ''::character varying,
					CONSTRAINT services_pk PRIMARY KEY (id),
					CONSTRAINT services_un UNIQUE (path)
				);
			`)
			return err
		}, func(db migrations.DB) error {
			log.Println("dropping table role_model.services...")
			_, err := db.Exec(`
				DROP table role_model.services;
			`)
			return err
		})
}
