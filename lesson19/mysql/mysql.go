package mysql

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql" // mysql database lib
)

// DbConfig data structure
type DbConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DbName   string
}

// DB type is
type DB struct {
	SQL *sql.DB
}

const dbConfigFile = "dbconfig.json"

var dbConfig = DbConfig{}

func init() {
	readConfig()
}

// readConfig reads a file that contains configuration parameters in JSON format
func readConfig() {
	file, err := os.Open(dbConfigFile)
	if err != nil {
		log.Fatalf("mysqldb: Cannot open config file[%s] - %v\n", dbConfigFile, err)
		return
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&dbConfig)
	if err != nil {
		log.Fatalf("mysqldb: Cannot get configuration from file[%s] - %v\n", dbConfigFile, err)
		return
	}
	//PrintJSON(dbConfig)
}

func (db *DB) ConnectMySQL() error {
	// compose mysql connection params string
	connParams := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		dbConfig.User, dbConfig.Password,
		dbConfig.Host, dbConfig.Port,
		dbConfig.DbName)
	//fmt.Println("MySQL:", connParams)
	conn, err := sql.Open("mysql", connParams)
	if err != nil {
		log.Fatalf("mysqldb: Cannot connect to MySQL database: %v\n", err)
	}
	db.SQL = conn
	return err
}

func PrintJSON(jsonObj interface{}) {
	s, _ := json.MarshalIndent(jsonObj, "", "\t")
	fmt.Println(string(s))
}
