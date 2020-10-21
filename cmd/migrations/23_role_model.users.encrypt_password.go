package main

import (
	"log"

	"github.com/go-pg/migrations/v8"
)

func init() {
	migrations.MustRegisterTx(
		func(db migrations.DB) error {
			log.Println("creating function encrypt_password()...")
			_, err := db.Exec(`
CREATE OR REPLACE FUNCTION role_model.encrypt_password()
 RETURNS trigger
 LANGUAGE plpgsql
AS $function$
	begin
		IF TG_OP = 'INSERT' then
			new."salt_pass" = public.gen_salt('bf'::text);
			new."password" = public.crypt(new."password", new."salt_pass");
		else
			return null;
		end if;
    return new;
END;
$function$
;
			`)
			return err
		}, func(db migrations.DB) error {
			log.Println("dropping function encrypt_password()...")
			_, err := db.Exec(`DROP FUNCTION IF EXISTS encrypt_password()`)
			return err
		})
}
