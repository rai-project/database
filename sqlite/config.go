package sqlite

import (
	"github.com/k0kubun/pp"
	"github.com/rai-project/config"
	"github.com/rai-project/database"
	"github.com/rai-project/vipertags"
)

type sqlitedbConfig struct {
	Provider       string        `json:"provider" config:"database.provider"`
	Endpoints      []string      `json:"endpoints" config:"database.endpoints"`
	MaxConnections int           `json:"max_connections" config:"database.max_connections" default:"0"`
	done           chan struct{} `json:"-" config:"-"`
}

// Config ...
var (
	Config = &sqlitedbConfig{
		done: make(chan struct{}),
	}
)

// ConfigName ...
func (sqlitedbConfig) ConfigName() string {
	return "SQLite"
}

// SetDefaults ...
func (a *sqlitedbConfig) SetDefaults() {
	vipertags.SetDefaults(a)
}

// Read ...
func (a *sqlitedbConfig) Read() {
	defer close(a.done)
	vipertags.Fill(a)
	if a.MaxConnections == 0 {
		a.MaxConnections = database.DefaultMaxConnections
	}
}

// Wait ...
func (c sqlitedbConfig) Wait() {
	<-c.done
}

// String ...
func (c sqlitedbConfig) String() string {
	return pp.Sprintln(c)
}

// Debug ...
func (c sqlitedbConfig) Debug() {
	log.Debug("SQLite Config = ", c)
}

func init() {
	config.Register(Config)
}
