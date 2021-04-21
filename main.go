package main

import (
	"context"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strconv"

	"github.com/jackc/pgx"
)

type result struct {
	Id      string
	Name    string
	Snippet template.HTML
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	defer conn.Close(context.Background())

	http.HandleFunc("/details", func(rw http.ResponseWriter, r *http.Request) {
		id := r.FormValue("id")

		if val, err := strconv.Atoi(id); err != nil || val < 0 {
			http.Error(rw, "500 - Internal Server Error", 500)
			return
		}

		var name string
		var types []string
		var classification string
		var abilities []string

		err := conn.QueryRow(context.Background(), sqlGetDetails, id).Scan(&id, &name, &types, &classification, &abilities)
		if err != nil {
			fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
			http.Error(rw, "500 - Internal Server Error", 500)
			return
		}

		detailsPage := template.New("")
		detailsPage, err = detailsPage.Parse(detailsHtml)
		if err != nil {
			panic(err)
		}

		err = detailsPage.Execute(rw, map[string]interface{}{
			"Id":             id,
			"Name":           name,
			"Types":          types,
			"Classification": classification,
			"Abilities":      abilities,
		})
		if err != nil {
			panic(err)
		}
	})

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		search := r.FormValue("search")

		if search == "" {
			homePage := template.New("")
			homePage, err = homePage.Parse(homeHtml)
			if err != nil {
				panic(err)
			}

			err = homePage.Execute(rw, nil)
			if err != nil {
				panic(err)
			}

			return
		}

		if len(search) > 51 {
			search = search[:51]
		}

		rows, err := conn.Query(context.Background(), sqlGetResults, search)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Query failed: %v\n", err)
			http.Error(rw, "500 - Internal Server Error", 500)
			return
		}

		defer rows.Close()

		results := make([]result, 0, 10)
		for rows.Next() {
			var r result
			if err := rows.Scan(&r.Id, &r.Name, &r.Snippet); err != nil {
				http.Error(rw, "500 - Internal Server Error", 500)
				return
			}
			results = append(results, r)
		}

		resultsPage := template.New("")
		resultsPage, err = resultsPage.Parse(resultsHtml)
		if err != nil {
			panic(err)
		}

		err = resultsPage.Execute(rw, map[string]interface{}{
			"Results": results,
			"Search":  search,
		})
		if err != nil {
			panic(err)
		}
	})

	http.ListenAndServe(":"+port, nil)
}
