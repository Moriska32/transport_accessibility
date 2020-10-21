package main

import (
	"log"

	"github.com/go-pg/migrations/v8"
)

func init() {
	migrations.MustRegisterTx(
		func(db migrations.DB) error {
			log.Println("creating triggers for role_model.services...")
			_, err := db.Exec(
				`			
create trigger create_service_trigger after insert
on role_model.services
for each row execute procedure role_model.create_service();
			`)
			return err
		}, func(db migrations.DB) error {
			log.Println("dropping creating triggers for role_model.services...")
			_, err := db.Exec(`
			DROP TRIGGER create_service_trigger ON role_model.services;
			`)
			return err
		})
}
