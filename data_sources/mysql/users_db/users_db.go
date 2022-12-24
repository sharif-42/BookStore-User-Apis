package users_db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	// import the package but we are not directly import in our code so a hyphen
	_ "github.com/go-sql-driver/mysql"
)

const (
	mysql_db_username = "mysql_db_username"
	mysql_db_password = "mysql_db_password"
	mysql_db_host     = "mysql_db_host"
	mysql_schema_name = "mysql_schema_name"
)

var (
	Client *sql.DB

	// load .env file
	err = godotenv.Load(".env")
	// TODO: need to handle env file import error

	// if err != nil {
	// log.Fatalf("Error loading .env file")
	// }

	USERNAME    = os.Getenv(mysql_db_username)
	PASSWORD    = os.Getenv(mysql_db_password)
	HOST        = os.Getenv(mysql_db_host)
	SCHEMA_NAME = os.Getenv(mysql_schema_name)
)

func init() {
	// <username>/<password>@tcp<host>/<schema_name>?charset=utf8
	log.Println("Connecting to Database.......")
	// fmt.Println(USERNAME, PASSWORD, HOST, SCHEMA_NAME)

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		USERNAME, PASSWORD, HOST, SCHEMA_NAME,
	)
	log.Println(dataSourceName)

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
