package main

import (
	"log"

	"github.com/go-pg/migrations/v8"
)

func init() {
	migrations.MustRegisterTx(
		func(db migrations.DB) error {
			log.Println("creating function role_model.update_user_role()...")
			_, err := db.Exec(`
CREATE OR REPLACE FUNCTION role_model.update_user_role()
 RETURNS trigger
 LANGUAGE plpgsql
AS $function$
   BEGIN
        IF (TG_OP = 'UPDATE') then
        	if new.role_id <> old.role_id then
        		update role_model.policies set policy_role_id = new.role_id where user_id = old.guid;
        	end if;
        END IF;
        RETURN NEW;
    END;
$function$
;
			`)
			return err
		}, func(db migrations.DB) error {
			log.Println("dropping function role_model.update_user_role()...")
			_, err := db.Exec(`DROP FUNCTION IF EXISTS role_model.update_user_role()`)
			return err
		})
}
