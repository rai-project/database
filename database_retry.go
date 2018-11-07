package database

// DO NOT EDIT!
// This code is generated with http://github.com/hexdigest/gowrap tool
// using https://raw.githubusercontent.com/hexdigest/gowrap/bd05dcaf6963696b62ac150a98a59674456c6c53/templates/retry template

//go:generate gowrap gen -d . -i Database -t https://raw.githubusercontent.com/hexdigest/gowrap/bd05dcaf6963696b62ac150a98a59674456c6c53/templates/retry -o database_retry.go

import (
	"time"
)

// DatabaseWithRetry implements Database interface instrumented with retries
type DatabaseWithRetry struct {
	Database
	_retryCount    int
	_retryInterval time.Duration
}

// NewDatabaseWithRetry returns DatabaseWithRetry
func NewDatabaseWithRetry(base Database, retryCount int, retryInterval time.Duration) DatabaseWithRetry {
	return DatabaseWithRetry{
		Database:       base,
		_retryCount:    retryCount + 1,
		_retryInterval: retryInterval,
	}
}

// Close implements Database
func (_d DatabaseWithRetry) Close() (err error) {
	for _i := 0; _i < _d._retryCount; _i++ {
		err = _d.Database.Close()
		if err == nil {
			break
		}
		if _d._retryCount > 1 {
			time.Sleep(_d._retryInterval)
		}
	}
	return err
}
