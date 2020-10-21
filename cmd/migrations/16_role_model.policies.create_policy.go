package main

import (
	"log"

	"github.com/go-pg/migrations/v8"
)

func init() {
	migrations.MustRegisterTx(
		func(db migrations.DB) error {
			log.Println("creating function role_model.create_policy()...")
			_, err := db.Exec(`
CREATE OR REPLACE FUNCTION role_model.create_policy()
 RETURNS trigger
 LANGUAGE plpgsql
AS $function$
  begin 
	PERFORM pg_notify(
		'add_new_policy',
		json_build_object(
		'operation', TG_OP,
		'record', row_to_json(NEW)
		)::text
	);
  return new;        
  end;
$function$
;
			`)
			return err
		}, func(db migrations.DB) error {
			log.Println("dropping function role_model.create_policy()...")
			_, err := db.Exec(`DROP FUNCTION IF EXISTS role_model.create_policy()`)
			return err
		})
}
