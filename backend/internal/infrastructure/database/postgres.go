package database

import (
	"database/sql"
	"fmt"
	"time"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func Connect(url string) (*sql.DB,error) {
	if url == "" { return nil, fmt.Errorf("DATABASE_URL no está configurada") }
	db,err := sql.Open("pgx",url); if err != nil { return nil,err }
	db.SetMaxOpenConns(10); db.SetMaxIdleConns(5); db.SetConnMaxLifetime(time.Hour)
	if err := db.Ping(); err != nil { db.Close(); return nil,err }
	return db,nil
}
