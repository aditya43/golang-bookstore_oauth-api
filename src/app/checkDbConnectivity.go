package app

import "github.com/aditya43/golang-bookstore_oauth-api/src/client/cassandra"

func CheckDBConnectivity() {
	session, err := cassandra.GetSession()
	if err != nil {
		panic(err)
	}
	session.Close()
}
