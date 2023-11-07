package client

import (
	"fmt"
	"os"

	"github.com/gocql/gocql"
)

var CASSANDRA_HOST string = os.Getenv("CASSANDRA_HOST")

type Client struct {
	session *gocql.Session
}

func (c *Client) Connect() error {
	cluster := gocql.NewCluster(CASSANDRA_HOST)
	sesion, err := cluster.CreateSession()
	if err != nil {
		return fmt.Errorf("error connecting to Cassandra: %s", err)
	}

	c.session = sesion
	return nil
}

func (c *Client) CreateKeyspace(keyspace string) error {
	cql := fmt.Sprintf(`CREATE KEYSPACE IF NOT EXISTS %s WITH replication = {'class': 'SimpleStrategy', 'replication_factor': '1'}`, keyspace)
	return c.session.Query(cql).Exec()
}

func (c *Client) CreateTable(keyspace string, table string) error {
	cql := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s.%s (name text, gender text, phone text, PRIMARY KEY (name))`, keyspace, table)
	return c.session.Query(cql).Exec()
}

func (c *Client) InsertFakeData(keyspace string, table string, numUsers int) error {
	users := GetFakeUsers(numUsers)
	cql := fmt.Sprintf(`INSERT INTO %s.%s (name, gender, phone) VALUES(?, ?, ?)`, keyspace, table)

	batch := c.session.NewBatch(gocql.LoggedBatch)
	for _, user := range users {
		batch.Query(cql, user.Name, user.Gender, user.Phone)
	}

	return c.session.ExecuteBatch(batch)
}
