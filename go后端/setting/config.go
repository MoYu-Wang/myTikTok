package setting

type MySQLConfig struct {
	Host         string
	User         string
	Password     string
	DB           string
	Port         int
	MaxOpenConns int
	MaxIdleConns int
}
