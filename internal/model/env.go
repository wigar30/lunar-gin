package model

type EnvConfigs struct {
	AppName string `mapstructure:"APP_NAME"`
	AppPort string `mapstructure:"APP_PORT"`

	GinMode string `mapstructure:"GIN_MODE"`

	DbConnection string `mapstructure:"DB_CONNECTION"`
	DbHost       string `mapstructure:"DB_HOST"`
	DbPort       string `mapstructure:"DB_PORT"`
	DbDatabase   string `mapstructure:"DB_DATABASE"`
	DbUsername   string `mapstructure:"DB_USERNAME"`
	DbPassword   string `mapstructure:"DB_PASSWORD"`

	LogLevel int32 `mapstructure:"LOG_LEVEL"`

	JwtSecret    string `mapstructure:"AUTH_JWT_SECRET"`
	JwtExpiredIn string `mapstructure:"AUTH_JWT_TOKEN_EXPIRES_IN"`
}
