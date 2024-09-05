package main

import (
	"fmt"
	"net/http"

	api "groupie/handlers"
)

func main() {
	http.HandleFunc("/locations/", api.LocationHandler)
	http.HandleFunc("/", api.ArtistsHandler)
	http.HandleFunc("/artist/", api.ArtistHandler)
	http.HandleFunc("/dates/", api.DateHandler)

	fmt.Println("server running")
	http.ListenAndServe(":3000", nil)
}
