package main

import (
	"log"

	"github.com/go-pg/migrations/v8"
)

func init() {
	migrations.MustRegisterTx(
		func(db migrations.DB) error {
			log.Println("creating function role_model.create_service()...")
			_, err := db.Exec(`
CREATE OR REPLACE FUNCTION role_model.create_service()
 RETURNS trigger
 LANGUAGE plpgsql
AS $function$
  begin 
    insert into role_model.policies(policy_type_id, policy_role_id, service_id, permission_id)
    values
       (1, 1, new.id, 1),
       (1, 2, new.id, 1),
       (1, 3, new.id, 2);
  return new;        
  end;
$function$
;
			`)
			return err
		}, func(db migrations.DB) error {
			log.Println("dropping function role_model.create_service()...")
			_, err := db.Exec(`DROP FUNCTION IF EXISTS role_model.create_service();`)
			return err
		})
}
