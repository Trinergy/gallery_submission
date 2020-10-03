package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "kennytest"
	password = "password"
	dbname   = "test_db"
)

// DB is a global variable for all handlers to access
var DB *sql.DB

// Index returns a JSON payload of Images
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	rows, _ := DB.Query("SELECT id,url,flagged,created_at,updated_at FROM images")
	defer rows.Close()

	var images []Image

	for rows.Next() {
		var id int
		var flagged bool
		var url string
		var createdAt, updatedAt time.Time

		err := rows.Scan(&id, &url, &flagged, &createdAt, &updatedAt)
		if err != nil {
			log.Fatal(err)
		}

		image := Image{
			ID:        id,
			Flagged:   flagged,
			URL:       url,
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
		}

		images = append(images, image)
	}

	header := w.Header()
	header.Set("Content-Type", "application/json")
	header.Set("Access-Control-Allow-Methods", header.Get("Allow"))
	header.Set("Access-Control-Allow-Origin", "http://localhost:3000")

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(images)
}

// UpdateImage updates attributes of an Image
func UpdateImage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	flagged := r.FormValue("flagged")

	sqlStatement := `UPDATE images
	SET flagged = $1
	WHERE id = $2;`

	_, err := DB.Exec(sqlStatement, flagged, id)
	if err != nil {
		log.Fatal(err)
	}

	header := w.Header()
	header.Set("Access-Control-Allow-Methods", header.Get("Allow"))
	header.Set("Access-Control-Allow-Origin", "http://localhost:3000")

	w.WriteHeader(http.StatusOK)
}

// PreflightHandler manual
func PreflightHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	header := w.Header()
	header.Set("Access-Control-Allow-Methods", "PUT")
	header.Set("Access-Control-Allow-Origin", "http://localhost:3000")

	w.WriteHeader(http.StatusNoContent)
}

func main() {
	// setup router
	router := httprouter.New()
	router.GET("/", Index)
	router.PUT("/image/:id", UpdateImage)
	router.OPTIONS("/image/:id", PreflightHandler)

	// establish db connection
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres",
		psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	DB = db

	// verify db connection
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(http.ListenAndServe(":8080", router))
}
