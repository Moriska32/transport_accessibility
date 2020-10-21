package main

import (
	"log"

	"github.com/go-pg/migrations/v8"
)

func init() {
	migrations.MustRegisterTx(
		func(db migrations.DB) error {
			log.Println("creating triggers for role_model.policies...")
			_, err := db.Exec(
				`			
create trigger policy_creation_trigger after insert
on role_model.policies
for each row execute procedure role_model.create_policy();

create trigger policy_update_trigger after update
on role_model.policies
for each row execute procedure role_model.update_policy();
	
			`)
			return err
		}, func(db migrations.DB) error {
			log.Println("dropping creating triggers for role_model.policies...")
			_, err := db.Exec(`
			DROP TRIGGER policy_creation_trigger ON role_model.policies;
			DROP TRIGGER policy_update_trigger ON role_model.policies;
			`)
			return err
		})
}
