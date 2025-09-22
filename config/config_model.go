package config

// ApplicationConfig 应用配置
type ApplicationConfig struct {
	Server   Server   `json:"server" yaml:"server"`
	Database Database `json:"database" yaml:"database"`
	Redis    Redis    `json:"redis" yaml:"redis"`
	Jwt      Jwt      `json:"jwt" yaml:"jwt"`
}

// Server 服务配置
type Server struct {
	LogPath  string `json:"logPath" yaml:"logPath"`
	LogLevel string `json:"logLevel" yaml:"logLevel"`
	Port     int    `json:"port" yaml:"port"`
	Name     string `json:"name" yaml:"name"`
}

// Database 数据库配置
type Database struct {
	IP       string `json:"ip" yaml:"ip"`
	Port     int    `json:"port" yaml:"port"`
	Username string `json:"username" yaml:"username"`
	Password string `json:"password" yaml:"password"`
	Database string `json:"database" yaml:"database"`
}

// Redis 缓存配置
type Redis struct {
	IP       string `json:"ip" yaml:"ip"`
	Port     int    `json:"port" yaml:"port"`
	Password string `json:"password" yaml:"password"`
	Database int    `json:"database" yaml:"database"`
}

// Jwt 令牌配置
type Jwt struct {
	SecretKey string `json:"secretKey" yaml:"secretKey"`
	Expire    int    `json:"expire" yaml:"expire"`
}
