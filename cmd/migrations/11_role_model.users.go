package main

import (
	"log"

	"github.com/go-pg/migrations/v8"
)

func init() {
	migrations.MustRegisterTx(
		func(db migrations.DB) error {
			log.Println("creating table role_model.users...")
			_, err := db.Exec(
				`
				CREATE TABLE role_model.users (
					guid uuid NOT NULL DEFAULT uuid_generate_v4(),
					login varchar NOT NULL,
					password varchar NULL,
					access varchar NOT NULL DEFAULT 'Authorization'::character varying,
					change_password bool NOT NULL DEFAULT true,
					nickname varchar NULL DEFAULT 'НЛО'::character varying,
					role_id int4 NOT NULL DEFAULT 3,
					register_tm timestamp NOT NULL DEFAULT now(),
					salt_pass varchar NOT NULL,
					CONSTRAINT users_pk PRIMARY KEY (guid),
					CONSTRAINT users_un UNIQUE (login)
				);
				CREATE INDEX users_idx ON role_model.users USING btree (guid);
			`)
			return err
		}, func(db migrations.DB) error {
			log.Println("dropping table role_model.users...")
			_, err := db.Exec(
				`DROP table role_model.users;
				`)
			return err
		})
}
