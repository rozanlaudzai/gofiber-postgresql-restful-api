package config

type Config struct {
	Server   *Server
	Database *Database
	Jwt      *Jwt
}

type Server struct {
	Host string
	Port string
}

type Database struct {
	Host     string
	Port     string
	Username string
	Password string
	TimeZone string
	Name     string
}

type Jwt struct {
	Key     string
	Expired int
}
