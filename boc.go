package main

import (
	"flag"
	"os"
	"database/sql"
	"github.com/bmizerany/pq"
	"fmt"
	"log"
	"net/http"
)

var query *string = flag.String("q", "select true", "Specify a query that retruns a bool.")
var pg *sql.DB

func main() {
	flag.Parse()

	cs, err := pq.ParseURL(os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Println("unable to parse database url")
		os.Exit(1)
	}
	cs += " sslmode=disable"
	pg, err = sql.Open("postgres", cs)
	if err != nil {
		fmt.Println("unable to connect to database")
		os.Exit(1)
	}

	http.HandleFunc("/check", func(w http.ResponseWriter, r *http.Request){
		if ok(pg, query) {
			w.WriteHeader(http.StatusOK) //200
			fmt.Fprintf(w, "OK")
		} else {
			w.WriteHeader(http.StatusExpectationFailed) //418
			fmt.Fprintf(w, "FAIL")
		}
	})
	port := ":"+os.Getenv("PORT")
	log.Printf("action=start port=%v", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

func ok(d *sql.DB, q *string) bool {
	log.Printf("query=%v", *q)
	var res bool
	var ok sql.NullBool
	d.QueryRow(*q).Scan(&ok)
	if ok.Valid {
		res = ok.Bool
	} else {
		res = false
	}
	log.Printf("result=%v", res)
	return res
}
