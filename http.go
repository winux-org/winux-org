package main

import "encoding/json"
import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type HttpHandler struct{}

func SetupHTTPRouting() {
	router := mux.NewRouter()
	router.HandleFunc("/", indexHandler)
	router.HandleFunc("/api/ideas", postsHandler)
	http.Handle("/", router)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!")
}

func postsHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	ideasSQL := "select * from ideas"
	rows, err := dbpool.Query(ctx, ideasSQL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "SQL query ended up with error : %v\n", err)
		os.Exit(1)
	}
	defer rows.Close()

	var results []Idea
	for rows.Next() {
		var idea Idea
		err = rows.Scan(&idea.Id, &idea.Email, &idea.Idea)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to scan %v\n", err)
			os.Exit(1)
		}
		results = append(results, idea)
	}
	if rows.Err() != nil {
		fmt.Fprintf(os.Stderr, "rows Error: %v\n", rows.Err())
		os.Exit(1)
	}

	jsonResponse, jsonError := json.Marshal(results)

	if jsonError != nil {
		fmt.Println("Unable to encode JSON")
	}

	//fmt.Println(string(jsonResponse))

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
