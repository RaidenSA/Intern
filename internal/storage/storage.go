package storage

import "sync"
import "Intern/internal/tokengen"

type MemoryStorage interface {
	//2 checks whether we have suck inquiry
	ValueToToken(string) (string, bool)
	TokenToValue(string) (string, bool)
	CreateToken(string) string
	//Create Token may be remade
}
type Container struct {
	mu              sync.Mutex
	MapTokenToValue map[string]string
	MapValueToToken map[string]string
}

func (c *Container) TokenToValue(token string) (string, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	//if _,ok:= c.mapTokenToValue[token]; ok{
	value, ok := c.MapTokenToValue[token]
	//}
	return value, ok
}

func (c *Container) ValueToToken(token string) (string, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	value, ok := c.MapValueToToken[token]
	return value, ok
}

func (c *Container) CreateToken(value string) string {
	c.mu.Lock()
	defer c.mu.Unlock()
	if token, ok := c.MapValueToToken[value]; ok {
		return token
	}
	token := tokengen.TokenGenerator()
	for _, ok := c.MapTokenToValue[token]; ok; {
		token = tokengen.TokenGenerator()
	}
	c.MapTokenToValue[token] = value
	c.MapValueToToken[value] = token
	return token
}
