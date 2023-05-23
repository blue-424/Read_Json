package repository

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"test/model"
)

func OpenConnection(obj model.Configuration) *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable", obj.Host, obj.Port, obj.User, obj.Password, obj.Dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return db
}
