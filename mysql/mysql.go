package mysql

import (
	_ "github.com/jinzhu/gorm/dialects/mysql" // include mysql dialect
	"github.com/rai-project/database"
	"github.com/rai-project/database/relational"
)

const (
	gormDialect = "myqsl"
)

func NewConnection(opts ...database.ConnectionOption) (database.Connection, error) {
	url := ""
	dbName := ""
	panic("not implemented...")
	return relational.NewConnection(gormDialect, url, dbName, opts...)
}
