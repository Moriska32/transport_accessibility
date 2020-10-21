package main

import (
	"log"

	"github.com/go-pg/migrations/v8"
)

func init() {
	migrations.MustRegisterTx(
		func(db migrations.DB) error {
			log.Println("creating table role_model.policies...")
			_, err := db.Exec(
				`			
				CREATE TABLE role_model.policies (
					id serial NOT NULL,
					policy_type_id int4 NOT NULL,
					policy_role_id int4 NOT NULL,
					user_id uuid NULL,
					service_id int4 NULL,
					permission_id int4 NULL,
					CONSTRAINT policies_permission_id_fkey FOREIGN KEY (permission_id) REFERENCES role_model.permissions(id),
					CONSTRAINT policies_policy_role_id_fkey FOREIGN KEY (policy_role_id) REFERENCES role_model.roles(id),
					CONSTRAINT policies_policy_type_id_fkey FOREIGN KEY (policy_type_id) REFERENCES role_model.policy_types(id),
					CONSTRAINT policies_service_id_fkey FOREIGN KEY (service_id) REFERENCES role_model.services(id) ON DELETE CASCADE,
					CONSTRAINT policies_user_id_fkey FOREIGN KEY (user_id) REFERENCES role_model.users(guid) ON DELETE CASCADE
				);
			`)
			return err
		}, func(db migrations.DB) error {
			log.Println("dropping table role_model.policies...")
			_, err := db.Exec(`
				DROP table role_model.policies; 
			`)
			return err
		})
}
