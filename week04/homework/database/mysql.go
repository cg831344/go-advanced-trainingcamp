package database

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	"sync"
	"time"
)

var DB *sql.DB
var once sync.Once

func Timeout(time time.Duration) func(*mysql.Config) {
	return func(c *mysql.Config) {
		c.Timeout = time
	}
}

func Conf(user, passwd, DBName string, options ...func(*mysql.Config)) *mysql.Config {
	c := &mysql.Config{
		User:   user,
		Passwd: passwd,
		DBName: DBName,
		Loc:    time.Local,
	}
	for _, option := range options {
		option(c)
	}
	return c
}

func InitDB(c *mysql.Config) {
	once.Do(func() {
		var err error
		DB, err = sql.Open("mysql", c.FormatDSN())
		if err != nil {
			panic(err)
		}
	})

}
