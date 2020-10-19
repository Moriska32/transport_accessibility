package configuration

// JWTKeys Информация о JWT ключах
type JWTKeys struct {
	PubFile string `json:"pub_file"`
	KeyFile string `json:"key_file"`
}
