package handler

import (
	"Intern/internal/storage"
	"fmt"
	"io"
	"log"
	"net/http"
)

const addr = "localhost:8080"

var InMemoryStorage = storage.Container{
	MapTokenToValue: make(map[string]string),
	MapValueToToken: make(map[string]string),
}

func TemplateHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		q := r.URL.String()
		// need to handle errors
		q = q[1:]
		if q == "" {
			http.Error(w, "No such url", http.StatusNotFound)
			return
		}
		//fmt.Println(q)
		if value, ok := InMemoryStorage.TokenToValue(q); ok {
			w.Header().Set("Content-Type", "text/plain; charset=utf-8")
			w.Header().Set("Location", value)
			w.WriteHeader(http.StatusTemporaryRedirect)
		} else {
			http.Error(w, "No such url", http.StatusNotFound)
		}
		return
	case http.MethodPost:
		longURL, err := io.ReadAll(r.Body)
		//fmt.Println(longURL)
		//fmt.Println(string(longURL))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		decodedURL := string(longURL)
		w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
		w.WriteHeader(http.StatusCreated)
		// If we already have this url, return token
		token := InMemoryStorage.CreateToken(decodedURL)
		//Return full short url
		_, err = fmt.Fprint(w, "http://"+addr+"/"+token)
		if err != nil {
			log.Fatal(err)
		}
		return
	default:
		http.Error(w, "wrong url", http.StatusBadRequest)
		return
	}
}
