package configuration

import (
	"encoding/json"
	"io/ioutil"
)

// AppConfiguration Конфигурация приложения
type AppConfiguration struct {
	PostgresConfiguration []*PGConf            `json:"postgres_database_cfg"`
	ServerCfg             *ServerConfiguration `json:"server_cfg"`
	JWTConfiguration      *JWTKeys             `json:"jwt_conf"`
	UseCORS               bool                 `json:"use_cors"`
	RBACFileName          string               `json:"rbac_conf"`
	DocsFolder            string               `json:"docs_folder"`
}

// NewConfiguration Инициализация конфигурации на основе содержимого файла конфигурации
func NewConfiguration(fname string) (*AppConfiguration, error) {
	configFile, err := ioutil.ReadFile(fname)
	if err != nil {
		return nil, err
	}
	cfg := AppConfiguration{}
	err = json.Unmarshal(configFile, &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
