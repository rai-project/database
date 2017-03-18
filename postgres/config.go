package postgres

import (
	"github.com/k0kubun/pp"
	"github.com/rai-project/config"
	"github.com/rai-project/database"
	"github.com/rai-project/vipertags"
)

type postgresdbConfig struct {
	Provider       string        `json:"provider" config:"database.provider"`
	Endpoints      []string      `json:"endpoints" config:"database.endpoints"`
	Username       string        `json:"username" config:"database.username"`
	Password       string        `json:"password" config:"database.password"`
	MaxConnections int           `json:"max_connections" config:"database.max_connections" default:"0"`
	Certificate    string        `json:"certificate" config:"database.certificate" default:""`
	DatabaseName   string        `json:"database_name" config:"database.database_name"`
	done           chan struct{} `json:"-" config:"-"`
}

var (
	Config = &postgresdbConfig{
		done: make(chan struct{}),
	}
)

func (postgresdbConfig) ConfigName() string {
	return "PostgreSQL"
}

func (a *postgresdbConfig) SetDefaults() {
	vipertags.SetDefaults(a)
}

func (a *postgresdbConfig) Read() {
	defer close(a.done)
	vipertags.Fill(a)
	if a.MaxConnections == 0 {
		a.MaxConnections = database.DefaultMaxConnections
	}
}

func (c postgresdbConfig) Wait() {
	<-c.done
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
