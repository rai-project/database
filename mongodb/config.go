package mongodb

import (
	"github.com/k0kubun/pp"
	"github.com/rai-project/config"
	"github.com/rai-project/database"
	"github.com/rai-project/vipertags"
)

type mongodbConfig struct {
	Provider        string        `json:"provider" config:"database.provider"`
	Endpoints       []string      `json:"endpoints" config:"database.endpoints"`
	Username        string        `json:"username" config:"database.username"`
	Password        string        `json:"password" config:"database.password"`
	Cert            string        `json:"cert" config:"database.cert"`
	InitialCapacity int           `json:"initial_capacity" config:"database.initial_capacity" default:"0"`
	MaxConnections  int           `json:"max_connections" config:"database.max_connections" default:"0"`
	done            chan struct{} `json:"-" config:"-"`
}

// Config ...
var (
	Config = &mongodbConfig{
		done: make(chan struct{}),
	}
)

// ConfigName ...
func (mongodbConfig) ConfigName() string {
	return "MongoDB"
}

// SetDefaults ...
func (a *mongodbConfig) SetDefaults() {
	vipertags.SetDefaults(a)
}

// Read ...
func (a *mongodbConfig) Read() {
	defer close(a.done)
	vipertags.Fill(a)
	if a.MaxConnections == 0 {
		a.MaxConnections = database.DefaultMaxConnections
	}
}

// Wait ...
func (c mongodbConfig) Wait() {
	<-c.done
}

// String ...
func (c mongodbConfig) String() string {
	return pp.Sprintln(c)
}

// Debug ...
func (c mongodbConfig) Debug() {
	log.Debug("MongoDB Config = ", c)
}

func init() {
	config.Register(Config)
}
