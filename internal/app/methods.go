package app

import (
	"intern/internal/tokengen"
)

func (s *Server) CreateToken(value string) string {

	if token, ok := s.Storage.ValueToToken(value); ok {
		return token
	}
	token := tokengen.TokenGenerator()
	for _, ok := s.Storage.TokenToValue(token); ok; {
		token = tokengen.TokenGenerator()
	}
	s.Storage.SetToken(token, value)
	return token
}
