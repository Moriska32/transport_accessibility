package datastore

// RoleModelPolicy Ролевые политики
type RoleModelPolicy struct {
	tableName struct{} `pg:"role_model.policies,alias:t,discard_unknown_columns"`

	ID           int                  `pg:"id,notnull" json:"-"`
	PolicyTypeID int                  `pg:"policy_type_id,notnull" json:"-"`
	PolicyRoleID int                  `pg:"policy_role_id,notnull" json:"-"`
	UserID       *string              `pg:"user_id" json:"-"`
	ServiceID    *int                 `pg:"service_id" json:"-"`
	PermissionID *int                 `pg:"permission_id" json:"-"`
	Permission   *RoleModelPermission `pg:"fk:permission_id,rel:has-one" json:"-"`
	PolicyRole   *RoleModelRole       `pg:"fk:policy_role_id,rel:has-one" json:"-"`
	PolicyType   *RoleModelPolicyType `pg:"fk:policy_type_id,rel:has-one" json:"-"`
	Service      *RoleModelService    `pg:"fk:service_id,rel:has-one" json:"-"`
	User         *User                `pg:"fk:user_id,rel:has-one" json:"-"`
}

// RoleModelPermission Виды разрешений
type RoleModelPermission struct {
	tableName struct{} `pg:"role_model.permissions,alias:t,discard_unknown_columns"`
	ID        int      `pg:"id,pk" json:"-"`
	PermName  string   `pg:"perm_name,notnull" json:"-"`
}

// RoleModelPolicyType Виды политик
type RoleModelPolicyType struct {
	tableName  struct{} `pg:"role_model.policy_types,alias:t,discard_unknown_columns"`
	ID         int      `pg:"id,pk" json:"-"`
	PolicyName string   `pg:"policy_name,notnull" json:"-"`
}

// RoleModelRole Виды ролей
type RoleModelRole struct {
	tableName struct{} `pg:"role_model.roles,alias:t,discard_unknown_columns"`
	ID        int      `pg:"id,pk" json:"-"`
	RoleName  string   `pg:"role_name,notnull" json:"-"`
}

// RoleModelService Существующие сервисы
type RoleModelService struct {
	tableName      struct{} `pg:"role_model.services,alias:t,discard_unknown_columns"`
	ID             int      `pg:"id,pk" json:"-"`
	MainPath       string   `pg:"main_path,notnull" json:"-"`
	APIVersionPath string   `pg:"api_version_path,notnull" json:"-"`
	Path           string   `pg:"path,notnull" json:"-"`
}

// User Пользователь системы
type User struct {
	tableName struct{} `pg:"role_model.users,alias:t,discard_unknown_columns"`

	ID             string             `pg:"guid,pk" json:"-"`
	Login          string             `pg:"login,notnull" json:"-"`
	RoleID         int                `pg:"role_id" json:"-"`
	Password       *string            `pg:"password" json:"-"`
	Access         string             `pg:"access,notnull" json:"-"`
	ChangePassword bool               `pg:"change_password,notnull" json:"-"`
	Nickname       *string            `pg:"nickname" json:"-"`
	PolicyRole     *UserRoleModelRole `pg:"fk:role_id,rel:has-one" json:"-"`
}

// UserRoleModelRole Виды ролей
type UserRoleModelRole struct {
	tableName struct{} `pg:"role_model.roles,alias:t,discard_unknown_columns"`
	ID        int      `pg:"id,pk" json:"-"`
	RoleName  string   `pg:"role_name,notnull" json:"-"`
}

// GetAllPolicies Набор всех ролевых политик
func (dbConn *DataBase) GetAllPolicies() ([]*RoleModelPolicy, error) {
	model := []*RoleModelPolicy{}
	err := dbConn.Model(&model).
		Relation("PolicyType.policy_name").
		Relation("PolicyRole.role_name").
		Relation("User.guid").
		Relation("Service.main_path").
		Relation("Service.api_version_path").
		Relation("Service.path").
		Relation("Permission.perm_name").
		Select()
	return model, err
}
