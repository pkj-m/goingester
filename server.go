package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

// handlerFunc handles incoming HTTP requests for /ingest endpoint
func handlerFunc(w http.ResponseWriter, r *http.Request) {
	// Message here to the response writer.
	fmt.Println("handlerFunc called")
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}
	fmt.Println("body:", string(body))
	err = ingestEvent(string(body))
	if err != nil {
		log.Printf("error in writing to DB: %s", err.Error())
		http.Error(w, "can't write to DB", http.StatusInternalServerError)
		return
	}
	log.Printf("event saved!")
	return
}

func ingestEvent(msg string) error {
	sqlStatement := `
INSERT INTO events (body)
VALUES ($1)`
	_, err := db.Exec(sqlStatement, msg)
	if err != nil {
		panic(err)
	}
	return nil
}
