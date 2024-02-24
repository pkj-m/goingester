package main

import (
	"database/sql"
	"fmt"
)

func ConnectDB() error {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=require",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)

	var err error
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		return err
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			fmt.Println("error in closing db")
		}
	}(db)

	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}
