package database

type Connection interface {
	Options() ConnectionOptions
	Close() error
}

type Database interface {
	Name() string
	Create() error
	Delete() error
}

type Table interface {
	Name() string
	Create() error
	Delete() error
	Insert(elem interface{}) error
}
