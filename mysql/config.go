package mysql

import (
	"github.com/k0kubun/pp"
	"github.com/rai-project/config"
	"github.com/rai-project/database"
	"github.com/rai-project/vipertags"
)

type mysqldbConfig struct {
	Provider       string   `json:"provider" config:"database.provider" default:"mysql"`
	Endpoints      []string `json:"endpoints" config:"database.endpoints"`
	Username       string   `json:"username" config:"database.username"`
	Password       string   `json:"password" config:"database.password"`
	MaxConnections int      `json:"max_connections" config:"database.max_connections" default:"0"`
	Certificate    string   `json:"certificate" config:"database.certificate" default:""`
	DatabaseName   string   `json:"database_name" config:"database.database_name"`
}

var (
	Config = &mysqldbConfig{}
)

func (mysqldbConfig) ConfigName() string {
	return "MySQL"
}

func (mysqldbConfig) SetDefaults() {
}

func (a *mysqldbConfig) Read() {
	vipertags.Fill(a)
	if a.MaxConnections == 0 {
		a.MaxConnections = database.DefaultMaxConnections
	}
}

func (c mysqldbConfig) String() string {
	return pp.Sprintln(c)
}

func (c mysqldbConfig) Debug() {
	log.Debug("MySQL Config = ", c)
}

func init() {
	config.Register(Config)
}
