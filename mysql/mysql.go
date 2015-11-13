package mysql

import (
	"database/sql"
	"fmt"

	"github.com/Sirupsen/logrus"
	_ "github.com/go-sql-driver/mysql"
)

const (
	openTemplate = "%s:%s@tcp(%s)/%s"
)

type (

	// MySQL have iformation for mySQL
	MySQL struct {
		isConnected bool
		db          *sql.DB
		config      *Config
	}

	// Config is the configuration for Mongo initiarization config option
	Config struct {
		Host     string
		Database string
		User     string
		Password string
	}

	// Connection is mySQL connection wrapper
	Connection struct {
		*sql.DB
	}
)

// NewDB creates new db
func NewDB(mc Config) (*MySQL, error) {
	m := &MySQL{
		config: &mc,
	}

	err := m.open()
	if err != nil {
		return nil, err
	}

	return m, nil
}

func (m *MySQL) open() error {
	if m.isConnected {
		return nil
	}

	logrus.WithFields(logrus.Fields{
		"host": m.config.Host,
		"db":   m.config.Database,
	}).Info("Connecting to MySQL")

	db, err := sql.Open("mysql", fmt.Sprintf(
		openTemplate,
		m.config.User,
		m.config.Password,
		m.config.Host,
		m.config.Database,
	))
	if err != nil {
		return err
	}

	m.db = db
	m.isConnected = true
	return nil
}

// GetConn gets connection to MySQL
func (m *MySQL) GetConn() (*Connection, error) {
	if !m.isConnected {
		err := m.open()
		if err != nil {
			return nil, err
		}
	}

	return &Connection{m.db}, nil
}
