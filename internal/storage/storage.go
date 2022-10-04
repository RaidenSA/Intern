package storage

import (
	"sync"
)

type MemoryStorage interface {
	//2 checks whether we have such inquiry
	ValueToToken(string) (string, bool) // Value, returns Token if found
	TokenToValue(string) (string, bool) // Token, returns Value if found
	SetToken(string, string)            // Token + value, places them in memory
	//Create Token may be remade
}
type Container struct {
	mu              sync.Mutex
	MapTokenToValue map[string]string
	MapValueToToken map[string]string
}
type DataBase struct {
	connStr string
}

func (db DataBase) TokenToValue(token string) (string, bool) {
	//open connection
	//do things (select)
	//close connection
	return "", false
}
func (db DataBase) ValueToToken(value string) (string, bool) {
	//open connection
	//do things (select)
	//close connection
	return "", false
}
func (db DataBase) SetToken(token string, value string) {
	//open connection
	//do things (select)-> insert
	//close connection
	return
}

func (c *Container) TokenToValue(token string) (string, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	//if _,ok:= c.mapTokenToValue[token]; ok{
	value, ok := c.MapTokenToValue[token]
	//}
	return value, ok
}

func (c *Container) ValueToToken(value string) (string, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	token, ok := c.MapValueToToken[value]
	return token, ok
}

func (c *Container) SetToken(token string, value string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.MapTokenToValue[token] = value
	c.MapValueToToken[value] = token
	return
}
