package storage

type MemoryStorage interface {
	//2 checks whether we have such inquiry
	ValueToToken(string) (string, bool) // Value, returns Token if found
	TokenToValue(string) (string, bool) // Token, returns Value if found
	SetToken(string, string)            // Token + value, places them in memory
	//Create Token may be remade
}
