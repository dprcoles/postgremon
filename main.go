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
	Id             int
	Name           string
	Types          [2]string
	Classification string
	Abilities      [2]string
	Snippet        template.HTML
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

	homePage := template.Must(template.New(".").Parse(homeHtml))
	resultsPage := template.Must(template.New(".").Parse(resultsHtml))
	detailsPage := template.Must(template.New(".").Parse(detailsHtml))

	http.HandleFunc("/details", func(w http.ResponseWriter, r *http.Request) {
		id := r.FormValue("id")

		if val, err := strconv.Atoi(id); err != nil || val < 0 {
			http.Error(w, "Not Found", 404)
			return
		}

		var name string
		var types [2]string
		var classification string
		var abilities [2]string

		err := conn.QueryRow(context.Background(), sqlGetDetails, id).Scan(&id, &name, &types, &classification, &abilities)
		if err != nil {
			fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
			http.Error(w, "Not Found", 404)
			return
		} else {
			detailsPage.Execute(w, map[string]interface{}{
				"Id":             id,
				"Name":           name,
				"Types":          types,
				"Classification": classification,
				"Abilities":      abilities,
			})
		}
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		query := r.FormValue("query")

		if query == "" {
			homePage.Execute(w, nil)
			return
		}

		if len(query) > 51 {
			query = query[:51]
		}

		rows, err := conn.Query(context.Background(), sqlGetResults, query)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Query failed: %v\n", err)
			http.Error(w, "Not Found", 404)
			return
		}

		defer rows.Close()

		results := make([]result, 0, 10)
		for rows.Next() {
			var r result
			if err := rows.Scan(&r.Id, &r.Name, &r.Snippet); err != nil {
				http.Error(w, "Not Found", 404)
				return
			}
			results = append(results, r)
		}

		if err := rows.Err(); err != nil {
			http.Error(w, "Not Found", 404)
			return
		}

		resultsPage.Execute(w, map[string]interface{}{
			"Results": results,
			"Query":   query,
		})
	})

	http.ListenAndServe(":"+port, nil)
}
