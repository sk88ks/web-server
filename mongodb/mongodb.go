package mongodb

import (
	"time"

	"github.com/Sirupsen/logrus"
	"gopkg.in/mgo.v2"
)

type (
	// Mongodb is the adapter for mongodb
	Mongodb struct {
		Connected bool
		DialInfo  *mgo.DialInfo
		Session   *mgo.Session
		Database  *mgo.Database
	}

	// Session is wrapper of mongodb session
	Session struct {
		*mgo.Session
		Database *mgo.Database
	}

	// Collection is wrapper of mongodb collection
	Collection struct {
		*mgo.Collection
	}

	// Config is the configuration for Mongo initiarization config option
	Config struct {
		Hosts    []string
		Timeout  time.Duration
		Database string
		User     string
		Password string
	}
)

// NewMongodb is creating new instance Mongodb
func NewMongodb(mc Config) (*Mongodb, error) {
	m := &Mongodb{}
	m.Connected = false
	m.DialInfo = &mgo.DialInfo{
		Addrs:    mc.Hosts,
		Database: mc.Database,
		Timeout:  mc.Timeout,
		Username: mc.User,
		Password: mc.Password,
	}

	err := m.connect()
	if err != nil {
		return m, err
	}

	return m, nil
}

// Create session to Mongodb server
// Dial is called just once for this process
func (m *Mongodb) connect() error {
	// When Connected is true, already Connected
	if m.Connected {
		return nil
	}

	logrus.WithFields(logrus.Fields{
		"hosts": m.DialInfo.Addrs,
		"db":    m.DialInfo.Database,
	}).Info("Connecting to mongodb")

	session, err := mgo.DialWithInfo(m.DialInfo)
	if err != nil {
		return err
	}

	session.SetMode(mgo.SecondaryPreferred, false)
	logrus.Info("Connected to mongod")

	m.Connected = true
	m.Session = session
	m.Database = session.DB(m.DialInfo.Database)
	return nil
}

// GetSession requests a socket connection from the session
// The connection is taken from connection pool
// When end to use connection, should Close the session
func (m *Mongodb) GetSession() (session *Session, err error) {

	if !m.Connected {
		if err = m.connect(); err != nil {
			return
		}
	}

	// Get from connection pool
	s := m.Session.Copy()
	return &Session{s, s.DB(m.Database.Name)}, nil
}

// GetCollection gets collection for a given collection name from Session
func (s *Session) GetCollection(collectionName string) (col *Collection) {
	return &Collection{s.Database.C(collectionName)}
}
