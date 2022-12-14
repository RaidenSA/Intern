package app

import (
	"context"
	"errors"
	"log"
	"strings"

	"intern/internal/api"
)

func (s *Server) Post(_ context.Context, req *api.PostRequest) (*api.PostResponse, error) {
	decodedURL := req.GetLongURL()
	if decodedURL == "" {
		log.Println("Empty Url")
		return nil, errors.New("EmptyUrl")
	}
	token := s.CreateToken(decodedURL)
	//Return full short url
	shortUrl := "http://" + Addr + "/" + token
	log.Println("GRPC POST request served. Got URL:", decodedURL, " Sent URL:", shortUrl)
	return &api.PostResponse{
		ShortURL: shortUrl,
	}, nil

}

func (s *Server) GET(_ context.Context, req *api.GetRequest) (*api.GetResponse, error) {
	q := req.GetShortURL()

	cutStr := "http://" + Addr + "/"
	q = strings.ReplaceAll(q, cutStr, "")
	if q == "" {
		log.Println("Empty Url")
		return nil, errors.New("EmptyUrl")
	}
	if value, ok := s.Storage.TokenToValue(q); ok {
		log.Println("GRP GET request served. Got token:", q, " Sent URL:", value)
		return &api.GetResponse{
			LongURL: value,
		}, nil
	} else {
		return nil, errors.New("404 No such url")
	}
}
