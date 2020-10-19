package datastore

import (
	"log"
	"strings"
)

// AData - Ролевая модель
type AData struct {
	tableName      struct{} `pg:"role_model.users,alias:t,discard_unknown_columns"`
	GUID           string   `pg:"guid" json:"-"`
	Login          string   `pg:"login" json:"username"`
	Password       string   `pg:"password" json:"-"`
	Access         string   `pg:"access" json:"-"`
	ChangePassword bool     `pg:"change_password" json:"-"`
	Nickname       string   `pg:"nickname" json:"-"`
	RoleID         string   `pg:"role_id" json:"-"`
	SudirID        int      `pg:"sudir_uid" json:"-"`
}

// GetAData Поиск пользователя по логину
func (dbConn *DataBase) GetAData(login string) (*AData, error) {
	model := AData{}
	err := dbConn.
		Model(&model).
		Column("guid", "login", "password", "access", "change_password", "nickname", "role_id").
		Where("lower(login)= ?", strings.ToLower(login)).
		Limit(1).
		Select()
	if err != nil {
		log.Println("ad writer", err)
		return nil, err
	}
	return &model, err
}

// GetADataPwd Поиск пользователя по логину и паролю
func (dbConn *DataBase) GetADataPwd(login, password string) (*AData, error) {
	model := AData{}
	err := dbConn.
		Model(&model).
		Column("guid", "login", "password", "access", "change_password", "nickname", "role_id").
		Where("lower(login) = ? and \"password\" = crypt(?, \"password\")", strings.ToLower(login), password).
		Limit(1).
		Select()
	if err != nil {
		log.Println("ad writer", err)
		return nil, err
	}
	return &model, err
}

// GetRole Роль пользователя
func (dbConn *DataBase) GetRole(userID string) (User, error) {
	model := User{}
	err := dbConn.Model(&model).
		Column("t.role_id").
		Relation("PolicyRole.role_name").
		Where("guid = ?", userID).
		Limit(1).
		Select()

	return model, err
}
