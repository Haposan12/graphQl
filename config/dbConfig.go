package config

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var GlobalDbSql *sql.DB

func NewConnection(conn string) (con *sql.DB, err error)  {
	con, err = sql.Open("postgres", conn)
	if err != nil{
		panic(err)
	}
	err = con.Ping()
	if err != nil {
		panic(err)
	}

	return
}

func StringConnection(host string, port int, user string, password string, dbName string) string  {
	return fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable", host, port, user, dbName, password)
}

func init()  {
	connInfo := StringConnection("localhost", 5432, "postgres", "postgres", "coba")
	GlobalDbSql, _ = NewConnection(connInfo)
}