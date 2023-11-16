package main

import (
	"cassandra-client/client"
)

const (
	keyspace = "test"
	table    = "users"
)

func main() {
	c := client.Client{}

	err := c.Connect()
	if err != nil {
		panic(err)
	}

	err = c.CreateKeyspace(keyspace)
	if err != nil {
		panic(err)
	}

	c.CreateTable(keyspace, table)
	if err != nil {
		panic(err)
	}

	c.WriteUsers(keyspace, table, 500, 10000)

	c.ReadUsers(keyspace, table, 100)
}
