package database

// DO NOT EDIT!
// This code is generated with http://github.com/hexdigest/gowrap tool
// using https://raw.githubusercontent.com/hexdigest/gowrap/bd05dcaf6963696b62ac150a98a59674456c6c53/templates/syncpool template

//go:generate gowrap gen -d . -i Table -t https://raw.githubusercontent.com/hexdigest/gowrap/bd05dcaf6963696b62ac150a98a59674456c6c53/templates/syncpool -o table_syncpool.go

// TablePool implements Table that uses pool of Table
type TablePool struct {
	pool chan Table
}

// NewTablePool takes several implementations of the Table and returns an instance of the Table
// that uses sync.Pool of given implemetations
func NewTablePool(impls ...Table) TablePool {
	if len(impls) == 0 {
		panic("empty pool")
	}

	pool := make(chan Table, len(impls))
	for _, i := range impls {
		pool <- i
	}

	return TablePool{pool: pool}
}

// Create implements Table
func (_d TablePool) Create(e interface{}) (err error) {
	_impl := <-_d.pool
	defer func() {
		_d.pool <- _impl
	}()
	return _impl.Create(e)
}

// Delete implements Table
func (_d TablePool) Delete() (err error) {
	_impl := <-_d.pool
	defer func() {
		_d.pool <- _impl
	}()
	return _impl.Delete()
}

// Insert implements Table
func (_d TablePool) Insert(elem interface{}) (err error) {
	_impl := <-_d.pool
	defer func() {
		_d.pool <- _impl
	}()
	return _impl.Insert(elem)
}

// Name implements Table
func (_d TablePool) Name() (s1 string) {
	_impl := <-_d.pool
	defer func() {
		_d.pool <- _impl
	}()
	return _impl.Name()
}
