package relational

import (
	"errors"

	"github.com/rai-project/database"
)

type relationalTable struct {
	db        *relationalDBConnection
	tableName string
}

func NewTable(conn database.Database, tableName string) (database.Table, error) {
	rconn, ok := conn.(*relationalDBConnection)
	if !ok {
		return nil, errors.New("Invalid database Database input. Expecting a relationaldb Database instance.")
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
