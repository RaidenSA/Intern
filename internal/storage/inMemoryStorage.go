package storage

import "sync"

type Container struct {
	mu              sync.Mutex
	MapTokenToValue map[string]string
	MapValueToToken map[string]string
}

func (c *Container) TokenToValue(token string) (string, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	value, ok := c.MapTokenToValue[token]
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
