package conf

type Redis struct {
	Host         string
	Port         string
	Password     string
	IdleTimeout  string
	ReadTimeout  string
	WriteTimeout string
	MinIdleConns int
	Retries      int
	DB           int
	PoolSize     int
}
