package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Masterminds/squirrel"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

var DB *sql.DB

func NewDB() (*sql.DB, error) {

	viper.AddConfigPath("./config")
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalln("cannot read from a config")
	}
	host := viper.Get("database.host").(string)
	port := viper.Get("database.port").(float64)
	user := viper.Get("database.user").(string)
	dbname := viper.Get("database.dbname").(string)
	password := viper.Get("database.password").(string)

	// Starting a database
	psqlInfo := fmt.Sprintf("postgresql://%s:%s@%s:%v/%s?sslmode=disable", user, password, host, port, dbname)
	DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}
	return DB, nil
}

func QueryBuilder(db *sql.DB) squirrel.StatementBuilderType {
	return squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).RunWith(db)
}
