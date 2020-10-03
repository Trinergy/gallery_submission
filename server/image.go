package main

import "time"

// Image contains information of location of image for render
type Image struct {
	ID        int       `json:"id"`
	URL       string    `json:"url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Flagged   bool      `json:"flagged"`
}

// Images is a collection of the Image struct
type Images []Image
