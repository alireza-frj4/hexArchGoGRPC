package db

import (
	"database/sql"
	"log"
	"time"

	sq "github.com/Masterminds/squirrel"
	_ "github.com/go-sql-driver/mysql"
)

type Adapter struct {
	db *sql.DB
}

func NewAdapter(driverName, dataSourceName string) (*Adapter, error) {
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		log.Fatalf("db connection failur: %v", err)
	}

	//test connection
	err = db.Ping()
	if err != nil {
		log.Fatalf("db Ping failur: %v", err)
	}
	return &Adapter{db: db}, nil
}

func (da Adapter) CloseDbConnecton() {
	err := da.db.Close()
	if err != nil {
		log.Fatalf("db close failur %v", err)
	}
}

func (da Adapter) AddToHistory(answer int32, operation string) error {
	queryString, args, err := sq.Insert("arithm_history").Columns("data", "answer", "operation").
		Values(time.Now(), answer, operation).ToSql()
	if err != nil {
		return err
	}

	_, err = da.db.Exec(queryString, args...)
	if err != nil {
		return nil
	}

	return nil
}
