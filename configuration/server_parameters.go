package configuration

// ServerConfiguration Информация о сервере
type ServerConfiguration struct {
	Port     string `json:"port"`
	APIPath  string `json:"api_path"`
	GRPCPort string `json:"grpc_port"`
}
