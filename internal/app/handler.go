package app

import (
	"fmt"
	_ "github.com/lib/pq"
	"io"
	"log"
	"net/http"
)

const Addr = "localhost:8080"

func (s *Server) Handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		q := r.URL.String()
		// need to handle errors
		q = q[1:]
		if q == "" {
			http.Error(w, "No such url", http.StatusNotFound)
			return
		}
		if value, ok := s.Storage.TokenToValue(q); ok {
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
		token := s.CreateToken(decodedURL)
		//Return full short url
		_, err = fmt.Fprint(w, "http://"+Addr+"/"+token)
		if err != nil {
			fmt.Println("!")
			log.Fatal(err, "!")
		}
		log.Println("finished successfully")
		return
	default:
		http.Error(w, "wrong url", http.StatusBadRequest)
		return
	}
}
