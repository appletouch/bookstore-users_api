package users_db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

//Go 1.10 or higher
//MySQL (4.1+), MariaDB, Percona Server, Google CloudSQL or Sphinx (2.2.3+)
//please run  go get -u github.com/go-sql-driver/mysql

//WARNING: voor postgres databases use: https://github.com/jackc/pgx

import (
	_ "github.com/go-sql-driver/mysql"
)

var (
	Client *sql.DB
)

const (
	mysql_username = "env_mysql_username"
	mysql_password = "env_mysql_password"
	mysql_host     = "env_mysql_host"
	mysql_shema    = "env_mysql_shema"
)

func init() {
	//envionment variable declared in the startup env/ u can also us secrets
	username := os.Getenv(mysql_username)
	password := os.Getenv(mysql_password)
	host := os.Getenv(mysql_host)
	shema := os.Getenv(mysql_shema)

	//open functon once and use it during live of application.
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", username, password, host, shema)
	var err error
	Client, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}
	if err = Client.Ping(); err != nil {
		panic(err)
	}
	//mysql.SetLogger()
	log.Println("Database is successfully configured")

}
