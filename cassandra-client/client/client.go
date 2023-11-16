package client

import (
	"cassandra-client/data"
	"fmt"
	"os"

	"github.com/gocql/gocql"
)

var CassandraHost, UserName, Password string

func init() {
	CassandraHost = os.Getenv("CASSANDRA_HOST")
	UserName = os.Getenv("CASSANDRA_USERNAME")
	Password = os.Getenv("CASSANDRA_PASSWORD")
	if CassandraHost == "" {
		CassandraHost = "localhost"
	}
	if UserName == "" {
		UserName = "cassandra"
	}
	if Password == "" {
		Password = "cassandra"
	}
}

type Client struct {
	session *gocql.Session
}

func (c *Client) Connect() error {

	cluster := gocql.NewCluster(CassandraHost)
	cluster.Authenticator = gocql.PasswordAuthenticator{
		Username: UserName,
		Password: Password,
	}
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

func (c *Client) ReadUsers(keyspace string, table string, limit int) ([]data.User, error) {
	cql := fmt.Sprintf("SELECT name, gender, phone FROM %s.%s LIMIT %d", keyspace, table, limit)
	iter := c.session.Query(cql).Iter()
	nowRows := iter.NumRows()

	users := make([]data.User, nowRows)

	idx := 0
	var user data.User
	for iter.Scan(&user.Name, &user.Gender, &user.Phone) {
		users[idx] = user
		idx++
	}

	if err := iter.Close(); err != nil {
		return nil, err
	}

	return users, nil
}

func (c *Client) WriteUsers(keyspace string, table string, batchSize int, numberOfUser int) error {
	users := data.CreateFakeUsers(numberOfUser)
	cql := fmt.Sprintf("INSERT INTO %s.%s (name, gender, phone) VALUES(?, ?, ?)", keyspace, table)

	batch := c.session.NewBatch(gocql.UnloggedBatch)
	for i, user := range users {
		batch.Query(cql, user.Name, user.Gender, user.Phone)

		if (i+1)%batchSize == 0 || i+1 == numberOfUser {
			if err := c.session.ExecuteBatch(batch); err != nil {
				return err
			}
			batch = c.session.NewBatch(gocql.LoggedBatch)
		}
	}
	return nil
}
