package cassandra

import (
	"github.com/gocql/gocql"
)

var (
	session *gocql.Session
)

func init() {
	// Connect to Cassandra cluster(not "database server")
	cluster := gocql.NewCluster("127.0.0.1") // localhost
	cluster.Keyspace = "oauth"
	cluster.Consistency = gocql.Quorum

	var err error
	if session, err = cluster.CreateSession(); err != nil {
		panic(err)
	}
}

func GetSession() *gocql.Session {
	return session
}