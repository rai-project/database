package database

// Database ...
type Database interface {
	Options() Options
	Session() interface{}
	Close() error
	String() string
}

// Table ...
type Table interface {
	Name() string
	Create(e interface{}) error
	Delete() error
	Insert(elem interface{}) error
}
