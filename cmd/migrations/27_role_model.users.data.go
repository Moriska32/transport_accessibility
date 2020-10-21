package main

import (
	"log"

	"github.com/go-pg/migrations/v8"
)

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		log.Println("Inserting data into role_model.users")
		_, err := db.Exec(`
		INSERT INTO role_model.users(guid, login, password, access, nickname, role_id, register_tm)
			VALUES('f5badc14-1b64-4a02-ab43-ff94e4462210', 'admin', 'admin', 'Authorization', 'admin', 1, '2019-01-01 00:00:00');
		INSERT INTO role_model.users (guid, login, password, access, nickname, role_id, register_tm)
			VALUES('0899d40c-f87c-4899-905a-d2dab536f211', 'user', 'user', 'Authorization', 'user', 2, '2019-01-01 00:00:00');
		INSERT INTO role_model.users (guid, login, password, access, nickname, role_id, register_tm)
			VALUES('5ce90d8e-31dd-4079-b01e-cffdd6fb6007', 'dev', 'dev', 'Authorization', 'dev', 3, '2019-01-01 00:00:00');
		`)
		return err
	}, func(db migrations.DB) error {
		log.Println("deleting all rows from role_model.users...")
		_, err := db.Exec(`delete from role_model.users`)
		return err
	})
}
