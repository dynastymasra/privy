package config

import "fmt"

type PostgresConfig struct {
	host        string
	port        int
	name        string
	username    string
	password    string
	maxOpenConn int
	maxIdleConn int
	logEnabled  bool
}

func (p PostgresConfig) Host() string {
	return p.host
}

func (p PostgresConfig) Port() int {
	return p.port
}

func (p PostgresConfig) Name() string {
	return p.name
}

func (p PostgresConfig) Username() string {
	return p.username
}

func (p PostgresConfig) Password() string {
	return p.password
}

func (p PostgresConfig) MaxOpenConn() int {
	return p.maxOpenConn
}

func (p PostgresConfig) MaxIdleConn() int {
	return p.maxIdleConn
}

func (p PostgresConfig) LogEnabled() bool {
	return p.logEnabled
}

func (p PostgresConfig) ConnectionString() string {
	return fmt.Sprintf("dbname=%s user=%s password='%s' host=%s port=%d sslmode=disable",
		p.name, p.username, p.password, p.host, p.port)
}
