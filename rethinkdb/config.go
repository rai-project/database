package rethinkdb

import (
	"github.com/k0kubun/pp"
	"github.com/rai-project/config"
	"github.com/rai-project/database"
	"github.com/rai-project/vipertags"
)

type rethinkdbConfig struct {
	Provider        string        `json:"provider" config:"database.provider"`
	Endpoints       []string      `json:"endpoints" config:"database.endpoints"`
	Username        string        `json:"username" config:"database.username"`
	Password        string        `json:"password" config:"database.password"`
	AuthKey         string        `json:"authkey" config:"database.authkey"`
	Cert            string        `json:"cert" config:"database.cert"`
	InitialCapacity int           `json:"initial_capacity" config:"database.initial_capacity" default:"0"`
	MaxConnections  int           `json:"max_connections" config:"database.max_connections" default:"0"`
	done            chan struct{} `json:"-" config:"-"`
}

// Config ...
var (
	Config = &rethinkdbConfig{
		done: make(chan struct{}),
	}
)

// ConfigName ...
func (rethinkdbConfig) ConfigName() string {
	return "RethinkDB"
}

// SetDefaults ...
func (a *rethinkdbConfig) SetDefaults() {
	vipertags.SetDefaults(a)
}

// Read ...
func (a *rethinkdbConfig) Read() {
	defer close(a.done)
	vipertags.Fill(a)
	if a.AuthKey == "" {
		a.AuthKey = a.Password
	}
	if a.InitialCapacity == 0 {
		a.InitialCapacity = DefaultInitialCapacity
	}
	if a.MaxConnections == 0 {
		a.MaxConnections = database.DefaultMaxConnections
	}
}

// Wait ...
func (c rethinkdbConfig) Wait() {
	<-c.done
}

// String ...
func (c rethinkdbConfig) String() string {
	return pp.Sprintln(c)
}

// Debug ...
func (c rethinkdbConfig) Debug() {
	log.Debug("RethinkDB Config = ", c)
}

func init() {
	config.Register(Config)
}
