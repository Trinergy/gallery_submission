package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

// Index returns a JSON payload of Images
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// TODO: Fetch Images from DB instead of hard-coded
	images := []Image{
		{ID: 1, URL: "https://picsum.photos/200/300", UpdatedAt: time.Now(), CreatedAt: time.Now(), Flagged: false},
		{ID: 2, URL: "https://picsum.photos/200/300", UpdatedAt: time.Now(), CreatedAt: time.Now(), Flagged: true},
		{ID: 3, URL: "https://picsum.photos/200/300", UpdatedAt: time.Now(), CreatedAt: time.Now(), Flagged: true},
		{ID: 4, URL: "https://picsum.photos/200/300", UpdatedAt: time.Now(), CreatedAt: time.Now(), Flagged: false},
	}

	header := w.Header()
	header.Set("Content-Type", "application/json")
	header.Set("Access-Control-Allow-Methods", header.Get("Allow"))
	header.Set("Access-Control-Allow-Origin", "*")

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(images)
}

// UpdateImage updates attributes of an Image
func UpdateImage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	router.POST("/hello/:name", UpdateImage)

	log.Fatal(http.ListenAndServe(":8080", router))
}
