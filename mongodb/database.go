package mongodb

import (
	"context"
	"errors"
	"strings"
	"time"

	"gopkg.in/mgo.v2"
	"upper.io/db.v3"
	"upper.io/db.v3/mongo"

	"github.com/rai-project/config"
	"github.com/rai-project/database"
	"github.com/rai-project/utils"
)

type mongoDatabase struct {
	session      db.Database
	databaseName string
	opts         database.Options
}

type debugLogger struct {
}

func (lg *debugLogger) Log(m *db.QueryStatus) {
	log.Printf("\n\t%s\n\n", strings.Replace(m.String(), "\n", "\n\t", -1))
}

func (lg *debugLogger) Output(calldepth int, s string) error {
	log.Printf("%s\n", s)
	return nil
}

// NewDatabase ...
func NewDatabase(databaseName string, opts ...database.Option) (database.Database, error) {
	options := database.Options{
		Endpoints:      Config.Endpoints,
		Username:       Config.Username,
		Password:       Config.Password,
		TLSConfig:      nil,
		MaxConnections: Config.MaxConnections,
		Context:        context.Background(),
	}

	for _, o := range opts {
		o(&options)
	}

	if len(options.Endpoints) == 0 {
		return nil, errors.New("no endpoints found")
	}

	decrypt := func(s string) string {
		if strings.HasPrefix(s, utils.CryptoHeader) && config.App.Secret != "" {
			if val, err := utils.DecryptStringBase64(config.App.Secret, s); err == nil {
				return val
			}
		}
		return s
	}
	sess := &mongo.Source{
		Settings: db.NewSettings(),
	}
	sess.Settings.SetConnMaxLifetime(5 * time.Hour)
	sess.Settings.SetMaxIdleConns(options.MaxConnections)
	sess.Settings.SetMaxOpenConns(options.MaxConnections)

	if debug {
		sess.Settings.SetLogging(true)
		sess.Settings.SetLogger(&debugLogger{})
		mgo.SetLogger(&debugLogger{})
		mgo.SetDebug(true)
	}

	connectionURL := mongo.ConnectionURL{
		User:     options.Username,
		Password: decrypt(options.Password),
		Host:     options.Endpoints[0],
		Database: databaseName,
	}

	err := sess.Open(connectionURL)
	if err != nil {
		return nil, err
	}

	return &mongoDatabase{
		session:      sess,
		databaseName: databaseName,
		opts:         options,
	}, nil
}

// Session ...
func (conn *mongoDatabase) Session() interface{} {
	return conn.session
}

// Options ...
func (conn *mongoDatabase) Options() database.Options {
	return conn.opts
}

// Close ...
func (conn *mongoDatabase) Close() error {
	return conn.session.Close()
}

// String ...
func (conn *mongoDatabase) String() string {
	return "mongodb"
}

func (s *mongoDatabase) WithCollection(collection string, f func(db.Collection) error) error {
	c := s.session.Collection(collection)
	return f(c)
}
