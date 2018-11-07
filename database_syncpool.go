package database

// DO NOT EDIT!
// This code is generated with http://github.com/hexdigest/gowrap tool
// using https://raw.githubusercontent.com/hexdigest/gowrap/bd05dcaf6963696b62ac150a98a59674456c6c53/templates/syncpool template

//go:generate gowrap gen -d . -i Database -t https://raw.githubusercontent.com/hexdigest/gowrap/bd05dcaf6963696b62ac150a98a59674456c6c53/templates/syncpool -o database_syncpool.go

// DatabasePool implements Database that uses pool of Database
type DatabasePool struct {
	pool chan Database
}

// NewDatabasePool takes several implementations of the Database and returns an instance of the Database
// that uses sync.Pool of given implemetations
func NewDatabasePool(impls ...Database) DatabasePool {
	if len(impls) == 0 {
		panic("empty pool")
	}

	pool := make(chan Database, len(impls))
	for _, i := range impls {
		pool <- i
	}

	return DatabasePool{pool: pool}
}

// Close implements Database
func (_d DatabasePool) Close() (err error) {
	_impl := <-_d.pool
	defer func() {
		_d.pool <- _impl
	}()
	return _impl.Close()
}

// Options implements Database
func (_d DatabasePool) Options() (o1 Options) {
	_impl := <-_d.pool
	defer func() {
		_d.pool <- _impl
	}()
	return _impl.Options()
}

// Session implements Database
func (_d DatabasePool) Session() (p1 interface{}) {
	_impl := <-_d.pool
	defer func() {
		_d.pool <- _impl
	}()
	return _impl.Session()
}

// String implements Database
func (_d DatabasePool) String() (s1 string) {
	_impl := <-_d.pool
	defer func() {
		_d.pool <- _impl
	}()
	return _impl.String()
}
