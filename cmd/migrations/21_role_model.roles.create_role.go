package main

import (
	"log"

	"github.com/go-pg/migrations/v8"
)

func init() {
	migrations.MustRegisterTx(
		func(db migrations.DB) error {
			log.Println("creating function role_model.create_role()...")
			_, err := db.Exec(`
CREATE OR REPLACE FUNCTION role_model.create_role()
 RETURNS trigger
 LANGUAGE plpgsql
AS $function$
  begin 
    insert into role_model.policies(policy_type_id, policy_role_id, service_id, permission_id)
		select 1, new.id, t.id, 1
	from role_model.services as t;
  return new;        
  end;
$function$
;
			`)
			return err
		}, func(db migrations.DB) error {
			log.Println("dropping function role_model.create_role()...")
			_, err := db.Exec(`DROP FUNCTION IF EXISTS role_model.create_role()`)
			return err
		})
}
