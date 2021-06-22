package conf

import (
	"github.com/google/cg831344/go-advanced-trainingcamp/week04/homework/database"
	"time"
)

type MysqlConf struct {
	User    string
	Passwd  string
	DBName  string
	timeout int
}

// 假设配置文件已经格式化

func GetMysqlConfig() MysqlConf {
	return MysqlConf{
		User:    "hk",
		Passwd:  "123456",
		DBName:  "test",
		timeout: 2,
	}
}

func MysqlInit(conf MysqlConf) {
	timeout := time.Second * time.Duration(conf.timeout)
	database.Conf(conf.User,
		conf.Passwd,
		conf.DBName,
		database.Timeout(timeout))
}
