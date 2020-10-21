package main

import (
	"log"

	"github.com/go-pg/migrations/v8"
)

func init() {
	migrations.MustRegisterTx(
		func(db migrations.DB) error {
			log.Println("creating function role_model.create_user()...")
			_, err := db.Exec(`
CREATE OR REPLACE FUNCTION role_model.create_user()
 RETURNS trigger
 LANGUAGE plpgsql
AS $function$
  begin 
    insert into role_model.policies(policy_type_id, policy_role_id, user_id)
    values(2, new.role_id, new.guid);
  return new;        
  end;
$function$
;
			`)
			return err
		}, func(db migrations.DB) error {
			log.Println("dropping function role_model.create_user()...")
			_, err := db.Exec(`DROP FUNCTION IF EXISTS role_model.create_user()`)
			return err
		})
}
