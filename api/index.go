package handler

import (
	"gitmirror/helpers"
	"log"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	services := [2]string{"github.com", "gitlab.com"}

	for _, service := range services {
		serviceURL := "https://" + service + path
		log.Println("Checking", serviceURL)
		if helpers.CheckService(serviceURL) {
			log.Println("Redirecting to", serviceURL)
			http.Redirect(w, r, serviceURL, http.StatusFound)
			return
		}
	}

	fileServer := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.StripPrefix(path, http.FileServer(http.Dir("static"))).ServeHTTP(w, r)
	})
	fileServer.ServeHTTP(w, r)

	
}