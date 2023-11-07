package main

import (
	"cassandra-client/client"
)

const (
	keyspace = "test"
	table    = "users"
)

func main() {
	client := client.Client{}

	err := client.Connect()
	if err != nil {
		panic(err)
	}

	err = client.CreateKeyspace(keyspace)
	if err != nil {
		panic(err)
	}

	client.CreateTable(keyspace, table)
	if err != nil {
		panic(err)
	}

	for i := 0; i < 10; i++ {
		err = client.InsertFakeData(keyspace, table, 100)
		if err != nil {
			panic(err)
		}
	}
}
