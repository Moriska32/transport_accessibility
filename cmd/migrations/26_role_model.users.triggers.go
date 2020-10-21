package main

import (
	"log"

	"github.com/go-pg/migrations/v8"
)

func init() {
	migrations.MustRegisterTx(
		func(db migrations.DB) error {
			log.Println("creating triggers for role_model.users...")
			_, err := db.Exec(`
create trigger new_user_trigger before insert
on role_model.users
for each row execute procedure role_model.encrypt_password();

create trigger user_creation_trigger after insert
on role_model.users
for each row execute procedure role_model.create_user();

create trigger user_upd_role_trigger after update
on role_model.users
for each row execute procedure role_model.update_user_role();
			`)
			return err
		}, func(db migrations.DB) error {
			log.Println("dropping creating triggers for role_model.users...")
			_, err := db.Exec("DROP TRIGGER new_user_trigger ON role_model.users;")
			if err != nil {
				return err
			}
			_, err = db.Exec("DROP TRIGGER user_creation_trigger ON role_model.users;")
			if err != nil {
				return err
			}
			_, err = db.Exec("DROP TRIGGER user_upd_role_trigger ON role_model.users;")
			return err
		})
}
