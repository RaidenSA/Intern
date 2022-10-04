package app

import "intern/internal/storage"

type MemoryStorage interface {
	//2 checks whether we have such inquiry
	ValueToToken(string) (string, bool) // Value, returns Token if found
	TokenToValue(string) (string, bool) // Token, returns Value if found
	CreateToken(string) string
	SetToken(string, string) // Token + value, places them in memory
	//Create Token may be remade
}

type Service struct {
	CurStorage MemoryStorage
}

var NewStorage = Service{}

func New() {
	NewStorage.CurStorage = &storage.Container{
		MapTokenToValue: make(map[string]string),
		MapValueToToken: make(map[string]string),
	}
	return
}
