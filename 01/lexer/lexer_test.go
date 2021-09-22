package lexer

import (
	"testing"
)

func TestNextToken(t *testing.T) {

	input := `=+(){},;`
	tests := struct {
		expectedType token.Tokentype
	}{}
}
