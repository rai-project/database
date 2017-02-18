package ql

import (
	"github.com/k0kubun/pp"
	"github.com/rai-project/config"
	"github.com/rai-project/database"
	"github.com/rai-project/vipertags"
)

type qldbConfig struct {
	Provider       string   `json:"provider" config:"database.provider" default:"ql"`
	Endpoints      []string `json:"endpoints" config:"database.endpoints"`
	MaxConnections int      `json:"max_connections" config:"database.max_connections" default:"0"`
}

var (
	Config = &qldbConfig{}
)

func (qldbConfig) ConfigName() string {
	return "ql"
}

func (qldbConfig) SetDefaults() {
}

func (a *qldbConfig) Read() {
	vipertags.Fill(a)
	if a.MaxConnections == 0 {
		a.MaxConnections = database.DefaultMaxConnections
	}
}

func (c qldbConfig) String() string {
	return pp.Sprintln(c)
}

func (c qldbConfig) Debug() {
	log.Debug("ql Config = ", c)
}

func init() {
	config.Register(Config)
}
