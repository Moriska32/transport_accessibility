package main

import (
	"log"

	"github.com/go-pg/migrations/v8"
)

func init() {
	migrations.MustRegisterTx(
		func(db migrations.DB) error {
			log.Println("creating triggers for role_model.roles...")
			_, err := db.Exec(
				`			
create trigger policy_creation_trigger after insert
on role_model.roles
for each row execute procedure role_model.create_role();
			`)
			return err
		}, func(db migrations.DB) error {
			log.Println("dropping creating triggers for role_model.roles...")
			_, err := db.Exec(`
			DROP TRIGGER policy_creation_trigger ON role_model.roles;
			`)
			return err
		})
}
