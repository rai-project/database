package database

// DO NOT EDIT!
// This code is generated with http://github.com/hexdigest/gowrap tool
// using https://raw.githubusercontent.com/hexdigest/gowrap/bd05dcaf6963696b62ac150a98a59674456c6c53/templates/opentracing template

//go:generate gowrap gen -d . -i Database -t https://raw.githubusercontent.com/hexdigest/gowrap/bd05dcaf6963696b62ac150a98a59674456c6c53/templates/opentracing -o database_tracing.go

// DatabaseWithTracing implements Database interface instrumented with opentracing spans
type DatabaseWithTracing struct {
	Database
	_instance string
}

// NewDatabaseWithTracing returns DatabaseWithTracing
func NewDatabaseWithTracing(base Database, instance string) DatabaseWithTracing {
	return DatabaseWithTracing{
		Database:  base,
		_instance: instance,
	}
}
