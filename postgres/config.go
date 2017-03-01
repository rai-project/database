package postgres

import (
	"github.com/k0kubun/pp"
	"github.com/rai-project/config"
	"github.com/rai-project/database"
	"github.com/rai-project/vipertags"
)

type postgresdbConfig struct {
	Provider       string   `json:"provider" config:"database.provider" default:"postgres"`
	Endpoints      []string `json:"endpoints" config:"database.endpoints"`
	Username       string   `json:"username" config:"database.username"`
	Password       string   `json:"password" config:"database.password"`
	MaxConnections int      `json:"max_connections" config:"database.max_connections" default:"0"`
	Certificate    string   `json:"certificate" config:"database.certificate" default:""`
	DatabaseName   string   `json:"database_name" config:"database.database_name"`
}

var (
	Config = &postgresdbConfig{}
)

func (postgresdbConfig) ConfigName() string {
	return "PostgreSQL"
}

func (postgresdbConfig) SetDefaults() {
}

func (a *postgresdbConfig) Read() {
	vipertags.Fill(a)
	if a.MaxConnections == 0 {
		a.MaxConnections = database.DefaultMaxConnections
	}
}

func (c postgresdbConfig) String() string {
	return pp.Sprintln(c)
}

func (c postgresdbConfig) Debug() {
	log.Debug("PostgreSQL Config = ", c)
}

func init() {
	config.Register(Config)
}
