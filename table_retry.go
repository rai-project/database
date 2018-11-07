package database

// DO NOT EDIT!
// This code is generated with http://github.com/hexdigest/gowrap tool
// using https://raw.githubusercontent.com/hexdigest/gowrap/bd05dcaf6963696b62ac150a98a59674456c6c53/templates/retry template

//go:generate gowrap gen -d . -i Table -t https://raw.githubusercontent.com/hexdigest/gowrap/bd05dcaf6963696b62ac150a98a59674456c6c53/templates/retry -o table_retry.go

import (
	"time"
)

// TableWithRetry implements Table interface instrumented with retries
type TableWithRetry struct {
	Table
	_retryCount    int
	_retryInterval time.Duration
}

// NewTableWithRetry returns TableWithRetry
func NewTableWithRetry(base Table, retryCount int, retryInterval time.Duration) TableWithRetry {
	return TableWithRetry{
		Table:          base,
		_retryCount:    retryCount + 1,
		_retryInterval: retryInterval,
	}
}

// Create implements Table
func (_d TableWithRetry) Create(e interface{}) (err error) {
	for _i := 0; _i < _d._retryCount; _i++ {
		err = _d.Table.Create(e)
		if err == nil {
			break
		}
		if _d._retryCount > 1 {
			time.Sleep(_d._retryInterval)
		}
	}
	return err
}

// Delete implements Table
func (_d TableWithRetry) Delete() (err error) {
	for _i := 0; _i < _d._retryCount; _i++ {
		err = _d.Table.Delete()
		if err == nil {
			break
		}
		if _d._retryCount > 1 {
			time.Sleep(_d._retryInterval)
		}
	}
	return err
}

// Insert implements Table
func (_d TableWithRetry) Insert(elem interface{}) (err error) {
	for _i := 0; _i < _d._retryCount; _i++ {
		err = _d.Table.Insert(elem)
		if err == nil {
			break
		}
		if _d._retryCount > 1 {
			time.Sleep(_d._retryInterval)
		}
	}
	return err
}
