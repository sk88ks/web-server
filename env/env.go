package env

import (
	"github.com/awa/liverpool-server/mongodb"
	"github.com/sk88ks/web-server/mysql"
)

// Environment expressed environment settings
type Environment struct {
	mySQL   *mysql.MySQL
	mongodb *mongodb.Mongodb
}

var env *Environment

func init() {
	env = &Environment{}

	mySQL, err := mysql.NewDB(mysql.Config{
		Host:     "host",
		Database: "dbname",
		User:     "user",
		Password: "pass",
	})
	if err != nil {
		panic(err)
	}
	env.mySQL = mySQL

	//mongo, err := mongodb.NewMongodb(mongodb.Config{
	//	Hosts:    []string{"127.0.0.1:27017"},
	//	Database: "",
	//	Timeout:  5 * time.Second,
	//})
	//if err != nil {
	//	panic(err)
	//}
	//env.mongodb = mongo

}

// GetMySQLConn gets a connection for MySQL from connection pooling
func GetMySQLConn() (*mysql.Connection, error) {
	return env.mySQL.GetConn()
}

// GetMongoSession gets a mongodb session from pool
func GetMongoSession() (*mongodb.Session, error) {
	return env.mongodb.GetSession()
}
