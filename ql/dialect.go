package ql

import (
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/cznic/ql"
	"github.com/jinzhu/gorm"
)

type qlDialect struct {
	db *sql.DB
	gorm.DefaultForeignKeyNamer
}

func (qlDialect) GetName() string {
	return "ql"
}

func (s *qlDialect) SetDB(db *sql.DB) {
	s.db = db
}

func (qlDialect) BindVar(i int) string {
	return "$$" // ?
}

func (qlDialect) Quote(key string) string {
	return fmt.Sprintf(`"%s"`, key)
}

func (s *qlDialect) DataTypeOf(field *gorm.StructField) string {
	dataValue, sqlType, _, additionalType := gorm.ParseFieldStructForDialect(field, s)

	if sqlType == "" {
		switch dataValue.Kind() {
		case reflect.Bool:
			sqlType = "bool"
		case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uintptr:
			sqlType = "int32"
		case reflect.Int, reflect.Int64, reflect.Uint64:
			sqlType = "int64"
		case reflect.Float64:
			sqlType = "float64"
		case reflect.Float32:
			sqlType = "float32"
		case reflect.String:
			sqlType = "string"
		case reflect.Struct:
			if _, ok := dataValue.Interface().(time.Time); ok {
				sqlType = "time"
			} else if _, ok := dataValue.Interface().(time.Duration); ok {
				sqlType = "duration"
			}
		default:
			if _, ok := dataValue.Interface().([]byte); ok {
				sqlType = "blob"
			}
		}
	}

	if sqlType == "" {
		panic(fmt.Sprintf("invalid sql type %s (%s) for ql", dataValue.Type().Name(), dataValue.Kind().String()))
	}

	if strings.TrimSpace(additionalType) == "" {
		return sqlType
	}
	return fmt.Sprintf("%v %v", sqlType, additionalType)
}

func (s qlDialect) HasIndex(tableName string, indexName string) bool {
	var count int
	s.db.QueryRow("SELECT count(*) FROM __Index WHERE TableName = ? AND Name = ? ", tableName, indexName).Scan(&count)
	return count > 0
}

func (s qlDialect) HasTable(tableName string) bool {
	var count int
	s.db.QueryRow("SELECT count(*) FROM __Table WHERE Name=?", tableName).Scan(&count)
	return count > 0
}

func (s qlDialect) HasColumn(tableName string, columnName string) bool {
	var count int
	s.db.QueryRow("SELECT count(*) FROM __Column WHERE TableName = ? AND Name = ?", tableName, columnName).Scan(&count)
	return count > 0
}

func (s qlDialect) RemoveIndex(tableName string, indexName string) error {
	_, err := s.db.Exec(fmt.Sprintf("DROP INDEX %v", indexName))
	return err
}

func (s qlDialect) HasForeignKey(tableName string, foreignKeyName string) bool {
	return false
}

func (s qlDialect) LimitAndOffsetSQL(limit, offset interface{}) (sql string) {
	if limit != nil {
		if parsedLimit, err := strconv.ParseInt(fmt.Sprint(limit), 0, 0); err == nil && parsedLimit > 0 {
			sql += fmt.Sprintf(" LIMIT %d", parsedLimit)
		}
	}
	if offset != nil {
		if parsedOffset, err := strconv.ParseInt(fmt.Sprint(offset), 0, 0); err == nil && parsedOffset > 0 {
			sql += fmt.Sprintf(" OFFSET %d", parsedOffset)
		}
	}
	return
}

func (qlDialect) SelectFromDummyTable() string {
	return ""
}

func (qlDialect) LastInsertIDReturningSuffix(tableName, columnName string) string {
	return ""
}

func (s qlDialect) CurrentDatabase() (name string) {
	d, ok := s.db.(*ql.DB)
	if !ok {
		return
	}
	name = d.Name()
}

func init() {
	gorm.RegisterDialect("ql", &qlDialect{})
}
