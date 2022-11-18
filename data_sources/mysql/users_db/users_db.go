package users_db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql" // import the package but we are not directly import in our code so a hyphen
)

const (
	mysql_db_username = "mysql_db_username"
	mysql_db_password = "mysql_db_password"
	mysql_db_host     = "mysql_db_host"
	mysql_schema_name = "mysql_schema_name"
)

var (
	Client *sql.DB

	// get data from env varialbles
	username    = os.Getenv(mysql_db_username)
	password    = os.Getenv(mysql_db_password)
	host        = os.Getenv(mysql_db_host)
	schema_name = os.Getenv(mysql_schema_name)
)

func init() {
	// <username>/<password>@tcp<host>/<schema_name>?charset=utf8
	fmt.Println("Connecting to Database.......")

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		"root",
		"s1302042",
		"127.0.0.1:3306",
		"users_db",
	)
	var err error
	Client, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}
	// to check is connection actually working.
	if err = Client.Ping(); err != nil {
		panic(err)
	}
	// TODO: we will configure it later, to see every log related to it.
	// mysql.SetLogger("")

	log.Println("Database Successfully Connected!")

}
