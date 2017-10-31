package dh

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"
)

const (
	MYSQL string = "mysql"
)

//* 默认60秒超时
var timeoutSecond time.Duration = 60

//* 连接
type Connection interface {
	getCid() int64                                                                       //获取cid
	getDB() *sql.DB                                                                      //获取db
	getUrl() string                                                                      //获取url
	getUserName() string                                                                 //获取用户名
	getPassword() string                                                                 //获取连接名称
	getDbName() string                                                                   //获取数据库名称
	connect(url string, port int, userName string, password string, dbName string) error //连接数据库
}

type connection struct {
	url      string
	userName string
	password string
	dbName   string
	db       *sql.DB
	cid      int64
}

func (c *connection) getUrl() string {
	return c.url
}

func (c *connection) getUserName() string {
	return c.userName
}

func (c *connection) getDbName() string {
	return c.dbName
}

func (c *connection) getPassword() string {
	return c.password
}

func (c *connection) getDB() *sql.DB {
	return c.db
}

func (c *connection) getCid() int64 {
	return c.cid
}

func (c *connection) connect(url string, port int, userName string, password string, dbName string) error {
	db, err := sql.Open(MYSQL, userName+":"+password+"@tcp("+url+":"+strconv.Itoa(port)+")/"+dbName + "?parseTime=true")
	if err != nil {
		return fmt.Errorf("connection faild: %s", err)
	}
	c.url = url
	c.userName = userName
	c.password = password
	c.dbName = dbName
	c.db = db
	c.cid = getCid()
	return nil
}

var conn Connection

func getDB() (*sql.DB, error) {
	if conn.getDB() == nil {
		return nil, fmt.Errorf("db uninitialized")
	}
	return conn.getDB(), nil
}
