package relational

import (
	"errors"

	"github.com/rai-project/database"
)

type relationalTable struct {
	db        *relationalDatabase
	tableName string
}

func NewTable(conn database.Connection, tableName string) (database.Table, error) {
	rconn, ok := conn.(*relationalDatabase)
	if !ok {
		return nil, errors.New("Invalid database connection input. Expecting a relationaldb connection instance.")
	}

	return &relationalTable{
		db:        rconn,
		tableName: tableName,
	}, nil
}

func (tbl *relationalTable) Name() string {
	return tbl.tableName
}

func (tbl *relationalTable) Create() error {
	return tbl.db.CreateTable(tbl.tableName).Error
}

func (tbl *relationalTable) Delete() error {
	return tbl.db.DropTableIfExists(tbl.tableName).Error
}

func (tbl *relationalTable) Insert(elem interface{}) error {
	return tbl.db.Model(tbl.tableName).Create(elem).Error
}
