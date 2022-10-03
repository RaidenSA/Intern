package main

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"sync"
	"time"
)

//Two maps to make search with ease, mutex to prevent collisisons
type Container struct {
	mu              sync.Mutex
	mapTokenToValue map[string]string
	mapValueToToken map[string]string
}

const addr = "localhost:8080"
const tokenLen = 10

var memoryStorage = Container{
	mapTokenToValue: make(map[string]string),
	mapValueToToken: make(map[string]string),
}

// Unique token generator
func tokenGenerator() string {
	rand.Seed(time.Now().UnixNano())
	var alphabet = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890_")
	res := make([]rune, tokenLen)
	for i := range res {
		res[i] = alphabet[rand.Intn(len(alphabet))]
	}
	return string(res)
}

func TemplateHandler(w http.ResponseWriter, r *http.Request) {
	var token string
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
		memoryStorage.mu.Lock()
		defer memoryStorage.mu.Unlock()
		if _, ok := memoryStorage.mapTokenToValue[q]; ok {
			w.Header().Set("Content-Type", "text/plain; charset=utf-8")
			w.Header().Set("Location", memoryStorage.mapTokenToValue[q])
			w.WriteHeader(http.StatusTemporaryRedirect)
		} else {
			http.Error(w, "No such url", http.StatusNotFound)
		}
		return
	case http.MethodPost:
		longURL, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		//Data comes in form of Encoded url.Value{} container with "url" key
		decodedData, err := url.ParseQuery(string(longURL))
		decodedURL := decodedData.Get("url")
		if decodedURL == "" {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
		w.WriteHeader(http.StatusCreated)
		// If we already have this url, return token
		memoryStorage.mu.Lock()
		defer memoryStorage.mu.Unlock()
		if val, ok := memoryStorage.mapValueToToken[decodedURL]; ok {
			token = val
		} else {
			//Else generate new one
			token = tokenGenerator()
			//Repeat until we select an unused one
			for _, ok = memoryStorage.mapTokenToValue[token]; ok; {
				token = tokenGenerator()
			}
			memoryStorage.mapTokenToValue[token] = decodedURL
			memoryStorage.mapValueToToken[decodedURL] = token
		}
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

func main() {
	// In-Memory storage
	//mapTokenToValue = make(map[string]string)
	//mapValueToToken = make(map[string]string)
	//Selection of handler
	http.HandleFunc("/", TemplateHandler)
	//Server Properties
	server := &http.Server{
		Addr: addr,
	}
	log.Fatal(server.ListenAndServe())

}
