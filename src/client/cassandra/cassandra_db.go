package cassandra

import (
	"fmt"

	"github.com/gocql/gocql"
)

func init() {
	// Connect to Cassandra cluster
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "oauth"
	session, err := cluster.CreateSession()
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to Cassandra cluster successfully..")
	defer session.Close()
}
