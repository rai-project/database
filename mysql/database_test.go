package mysql

import (
	"testing"

	"github.com/k0kubun/pp"
	"github.com/rai-project/config"
	"github.com/stretchr/testify/assert"
)

// XXXTestConnection ...
func XXXTestConnection(t *testing.T) {
	config.Init()
	db, err := NewDatabase("abduld")
	if !assert.NoError(t, err) {
		return
	}
	pp.Println(db)
}
