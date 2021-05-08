package app

import "github.com/aditya43/golang-bookstore_oauth-api/src/client/cassandra"

func CheckDBConnectivity() {
	session := cassandra.GetSession()
	if session == nil {
		panic("Failed to establish database connection")
	}
	session.Close()
}
