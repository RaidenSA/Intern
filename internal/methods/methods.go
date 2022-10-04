package methods

import "intern/internal/tokengen"
import "intern/internal/service"

func CreateToken(value string) string {
	if token, ok := service.ServiceStorage.CurStorage.ValueToToken(value); ok {
		return token
	}
	token := tokengen.TokenGenerator()
	for _, ok := service.ServiceStorage.CurStorage.TokenToValue(token); ok; {
		token = tokengen.TokenGenerator()
	}
	service.ServiceStorage.CurStorage.SetToken(token, value)
	return token
}
