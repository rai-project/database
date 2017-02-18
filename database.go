package database

type Database interface {
	Options() Options
	Close() error
	String() string
}

type Table interface {
	Name() string
	Create(e interface{}) error
	Delete() error
	Insert(elem interface{}) error
}
