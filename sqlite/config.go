package sqlite

import (
	"github.com/k0kubun/pp"
	"github.com/rai-project/config"
	"github.com/rai-project/database"
	"github.com/rai-project/vipertags"
)

type sqlitedbConfig struct {
	Provider       string   `json:"provider" config:"database.provider"`
	Endpoints      []string `json:"endpoints" config:"database.endpoints"`
	MaxConnections int      `json:"max_connections" config:"database.max_connections" default:"0"`
}

var (
	Config = &sqlitedbConfig{}
)

func (sqlitedbConfig) ConfigName() string {
	return "SQLite"
}

func (sqlitedbConfig) SetDefaults() {
}

func (a *sqlitedbConfig) Read() {
	vipertags.Fill(a)
	if a.MaxConnections == 0 {
		a.MaxConnections = database.DefaultMaxConnections
	}
}

func (c sqlitedbConfig) String() string {
	return pp.Sprintln(c)
}

func (c sqlitedbConfig) Debug() {
	log.Debug("SQLite Config = ", c)
}

func init() {
	config.Register(Config)
}
