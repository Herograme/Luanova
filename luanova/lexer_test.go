package luanova

import (
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `local function test(x: number, y: string): boolean
	if x ~= 10 and y == "hello" then
		x += 5
		y ..= " world"
		return true
	elseif x <= 20 or x >= 0 then
		-- this is a line comment
		-* this is a 
		block comment *-
		x -= 1
		x *= 2
		x /= 3
		x %= 4
		x = x ^ 2
		return false
	end

	local arr = {1, 2, 3}
	local dict = {
		[1] = "one",
		["two"] = 2
	}

	for i = 1, 10, 2 do
		if i < 5 then
			continue
		end
		if i > 8 then
			break
		end
	end

	while true do
		local x = 1 + 2 * 3 / 4 % 5
		if x then break end
	end

	local concat = "hello" .. "world" ... "variadic"

	-- Type system tokens
	type Point = {
		x: number,
		y: number
	}

	type Optional = number?
	type Callback = (string) -> boolean
	`

	tests := []struct {
		expectedType    int
		expectedLiteral string
	}{
		{Local, "local"},
		{Function, "function"},
		{Literal, "test"},
		{LParen, "("},
		{Literal, "x"},
		{Colom, ":"},
		{Literal, "number"},
		{Comma, ","},
		{Literal, "y"},
		{Colom, ":"},
		{Literal, "string"},
		{RParen, ")"},
		{Colom, ":"},
		{Literal, "boolean"},
		{If, "if"},
		{Literal, "x"},
		{NotEqual, "~="},
		{Literal, "10"},
		{And, "and"},
		{Literal, "y"},
		{Equal, "=="},
		{StringDelim, "hello"},
		{Then, "then"},
		{Literal, "x"},
		{PlusAssign, "+="},
		{Literal, "5"},
		{Literal, "y"},
		{Concat, ".."},
		{Assign, "="},
		{StringDelim, " world"},
		{Return, "return"},
		{True, "true"},
		{ElseIf, "elseif"},
		{Literal, "x"},
		{LessEqual, "<="},
		{Literal, "20"},
		{Or, "or"},
		{Literal, "x"},
		{GreaterEqual, ">="},
		{Literal, "0"},
		{Then, "then"},
		{Comment, "-- this is a line comment"},
		{CommentBlock, "\t-* this is a \n\t\tblock comment *-"},
		{Literal, "x"},
		{SubAssign, "-="},
		{Literal, "1"},
		{Literal, "x"},
		{MultiAssign, "*="},
		{Literal, "2"},
		{Literal, "x"},
		{DivAssign, "/="},
		{Literal, "3"},
		{Literal, "x"},
		{ModAssign, "%="},
		{Literal, "4"},
		{Literal, "x"},
		{Assign, "="},
		{Literal, "x"},
		{Po, "^"},
		{Literal, "2"},
		{Return, "return"},
		{False, "false"},
		{End, "end"},
		{Local, "local"},
		{Literal, "arr"},
		{Assign, "="},
		{LBrace, "{"},
		{Literal, "1"},
		{Comma, ","},
		{Literal, "2"},
		{Comma, ","},
		{Literal, "3"},
		{RBrace, "}"},
		{Local, "local"},
		{Literal, "dict"},
		{Assign, "="},
		{LBrace, "{"},
		{LBrack, "["},
		{Literal, "1"},
		{RBrack, "]"},
		{Assign, "="},
		{StringDelim, "one"},
		{Comma, ","},
		{LBrack, "["},
		{StringDelim, "two"},
		{RBrack, "]"},
		{Assign, "="},
		{Literal, "2"},
		{RBrace, "}"},
		{For, "for"},
		{Literal, "i"},
		{Assign, "="},
		{Literal, "1"},
		{Comma, ","},
		{Literal, "10"},
		{Comma, ","},
		{Literal, "2"},
		{Do, "do"},
		{If, "if"},
		{Literal, "i"},
		{Less, "<"},
		{Literal, "5"},
		{Then, "then"},
		{Continue, "continue"},
		{End, "end"},
		{If, "if"},
		{Literal, "i"},
		{Greater, ">"},
		{Literal, "8"},
		{Then, "then"},
		{Break, "break"},
		{End, "end"},
		{End, "end"},
		{While, "while"},
		{True, "true"},
		{Do, "do"},
		{Local, "local"},
		{Literal, "x"},
		{Assign, "="},
		{Literal, "1"},
		{Plus, "+"},
		{Literal, "2"},
		{Multi, "*"},
		{Literal, "3"},
		{Div, "/"},
		{Literal, "4"},
		{Mod, "%"},
		{Literal, "5"},
		{If, "if"},
		{Literal, "x"},
		{Then, "then"},
		{Break, "break"},
		{End, "end"},
		{End, "end"},
		{Local, "local"},
		{Literal, "concat"},
		{Assign, "="},
		{StringDelim, "hello"},
		{Concat, ".."},
		{StringDelim, "world"},
		{Dots, "..."},
		{StringDelim, "variadic"},
		{Comment, "-- Type system tokens"},
		{Literal, "type"},
		{Literal, "Point"},
		{Assign, "="},
		{LBrace, "{"},
		{Literal, "x"},
		{Colom, ":"},
		{Literal, "number"},
		{Comma, ","},
		{Literal, "y"},
		{Colom, ":"},
		{Literal, "number"},
		{RBrace, "}"},
		{Literal, "type"},
		{Literal, "Optional"},
		{Assign, "="},
		{Literal, "number"},
		{Question, "?"},
		{Literal, "type"},
		{Literal, "Callback"},
		{Assign, "="},
		{LParen, "("},
		{Literal, "string"},
		{RParen, ")"},
		{Arrow, "->"},
		{Literal, "boolean"},
		{EOF, ""},
	}

	l := NewLexer(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%s, got=%s (literal=%q)",
				i, TokenName(tt.expectedType), TokenName(tok.Type), tok.Literal)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}

func TestIllegalToken(t *testing.T) {
	input := "@"
	l := NewLexer(input)
	tok := l.NextToken()

	if tok.Type != Illegal {
		t.Errorf("Expected %s token, got %s", TokenName(Illegal), TokenName(tok.Type))
	}
}

func TestStringTokens(t *testing.T) {
	input := `"hello world" "test"`
	l := NewLexer(input)

	tests := []struct {
		expectedType    int
		expectedLiteral string
	}{
		{StringDelim, "hello world"},
		{StringDelim, "test"},
		{EOF, ""},
	}

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%s, got=%s",
				i, TokenName(tt.expectedType), TokenName(tok.Type))
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}

func TestNumberTokens(t *testing.T) {
	input := "123 456.789"
	l := NewLexer(input)

	tests := []struct {
		expectedType    int
		expectedLiteral string
	}{
		{Literal, "123"},
		{Literal, "456"},
		{Dot, "."},
		{Literal, "789"},
		{EOF, ""},
	}

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%s, got=%s",
				i, TokenName(tt.expectedType), TokenName(tok.Type))
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}

func TestCommentTokens(t *testing.T) {
	input := `-- line comment
	-* block comment *-`
	l := NewLexer(input)

	tests := []struct {
		expectedType int
	}{
		{Comment},
		{CommentBlock},
		{EOF},
	}

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%s, got=%s",
				i, TokenName(tt.expectedType), TokenName(tok.Type))
		}
	}
}

func TestOperatorPrecedence(t *testing.T) {
	input := "a += 1 - 2 * 3 / 4 % 5 ^ 6"
	l := NewLexer(input)

	tests := []struct {
		expectedType    int
		expectedLiteral string
	}{
		{Literal, "a"},
		{PlusAssign, "+="},
		{Literal, "1"},
		{Sub, "-"},
		{Literal, "2"},
		{Multi, "*"},
		{Literal, "3"},
		{Div, "/"},
		{Literal, "4"},
		{Mod, "%"},
		{Literal, "5"},
		{Po, "^"},
		{Literal, "6"},
		{EOF, ""},
	}

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%s, got=%s",
				i, TokenName(tt.expectedType), TokenName(tok.Type))
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}

// TestTokenName verifica se todos os tokens são corretamente convertidos para string
func TestTokenName(t *testing.T) {
	// Mapa de tokens e seus nomes esperados
	tokenNameMap := map[int]string{
		Illegal:      "Illegal",
		EOF:          "EOF",
		Comment:      "Comment",
		CommentBlock: "CommentBlock",
		True:         "True",
		False:        "False",
		Literal:      "Literal",
		Function:     "Function",
		Local:        "Local",
		If:           "If",
		ElseIf:       "ElseIf",
		Else:         "Else",
		While:        "While",
		For:          "For",
		End:          "End",
		Then:         "Then",
		Repeat:       "Repeat",
		Continue:     "Continue",
		Break:        "Break",
		In:           "In",
		Return:       "Return",
		Do:           "Do",
		Not:          "Not",
		NotEqual:     "NotEqual",
		Or:           "Or",
		Greater:      "Greater",
		Less:         "Less",
		GreaterEqual: "GreaterEqual",
		LessEqual:    "LessEqual",
		And:          "And",
		Assign:       "Assign",
		Equal:        "Equal",
		PlusAssign:   "PlusAssign",
		SubAssign:    "SubAssign",
		MultiAssign:  "MultiAssign",
		DivAssign:    "DivAssign",
		ModAssign:    "ModAssign",
		Declaration:  "Declaration",
		Plus:         "Plus",
		Sub:          "Sub",
		Multi:        "Multi",
		Div:          "Div",
		Mod:          "Mod",
		Po:           "Po",
		Concat:       "Concat",
		StringDelim:  "StringDelim",
		LParen:       "LParen",
		RParen:       "RParen",
		LBrace:       "LBrace",
		RBrace:       "RBrace",
		LBrack:       "LBrack",
		RBrack:       "RBrack",
		Comma:        "Comma",
		Semi:         "Semi",
		Colom:        "Colom",
		Dot:          "Dot",
		Dots:         "Dots",
		Arrow:        "Arrow",
		Question:     "Question",
		String:       "String",
		Int:          "Int",
		Bool:         "Bool",
		Float:        "Float",
		999:          "Unknown", // Teste para o caso default
	}

	for token, expectedName := range tokenNameMap {
		name := TokenName(token)
		if name != expectedName {
			t.Errorf("TokenName(%d) = %q, esperado %q", token, name, expectedName)
		}
	}
}

// TestEdgeCases testa casos extremos não cobertos pelos testes principais
func TestEdgeCases(t *testing.T) {
	// Teste para ponto e vírgula
	input1 := ";"
	l1 := NewLexer(input1)
	tok1 := l1.NextToken()
	if tok1.Type != Semi || tok1.Literal != ";" {
		t.Errorf("Expected %s token, got %s", TokenName(Semi), TokenName(tok1.Type))
	}

	// Teste para ~ (não seguido de =)
	input2 := "~"
	l2 := NewLexer(input2)
	tok2 := l2.NextToken()
	if tok2.Type != Illegal || tok2.Literal != "~" {
		t.Errorf("Expected %s token, got %s", TokenName(Illegal), TokenName(tok2.Type))
	}

	// Teste para peekChar retornando 0 no fim da string
	input3 := "a"
	l3 := NewLexer(input3)
	_ = l3.NextToken() // Consome 'a'
	if l3.peekChar() != 0 {
		t.Errorf("Expected peekChar() to return 0 at end of input")
	}

	// Teste para comentário de bloco sem fim
	input4 := "-* comentário sem fim"
	l4 := NewLexer(input4)
	tok4 := l4.NextToken()
	if tok4.Type != CommentBlock {
		t.Errorf("Expected %s token, got %s", TokenName(CommentBlock), TokenName(tok4.Type))
	}
}

// TestKeywords testa palavras-chave específicas que não foram cobertas
func TestKeywords(t *testing.T) {
	// Testa palavras-chave não cobertas pelos testes principais
	keywords := map[string]int{
		"else": Else,
		"not":  Not,
		"in":   In,
	}

	for keyword, expectedToken := range keywords {
		l := NewLexer(keyword)
		tok := l.NextToken()
		if tok.Type != expectedToken {
			t.Errorf("Expected %s token for %q, got %s",
				TokenName(expectedToken), keyword, TokenName(tok.Type))
		}
	}
}

// TestComplexTokens testa combinações de tokens mais complexas
func TestComplexTokens(t *testing.T) {
	input := `
	for i in pairs(table) do
		if not condition then
			print("not covered")
		else
			-- outro comentário
			break
		end
	end
	`

	l := NewLexer(input)

	// Verificamos se o lexer consegue processar esta entrada corretamente
	// sem falhar e garantindo que todos os tokens sejam reconhecidos
	tokens := []struct {
		expectedType    int
		expectedLiteral string
	}{
		{For, "for"},
		{Literal, "i"},
		{In, "in"},
		{Literal, "pairs"},
		{LParen, "("},
		{Literal, "table"},
		{RParen, ")"},
		{Do, "do"},
		{If, "if"},
		{Not, "not"},
		{Literal, "condition"},
		{Then, "then"},
		{Literal, "print"},
		{LParen, "("},
		{StringDelim, "not covered"},
		{RParen, ")"},
		{Else, "else"},
		{Comment, "-- outro comentário"},
		{Break, "break"},
		{End, "end"},
		{End, "end"},
		{EOF, ""},
	}

	for i, tt := range tokens {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%s, got=%s (literal=%q)",
				i, TokenName(tt.expectedType), TokenName(tok.Type), tok.Literal)
		}

		if tt.expectedLiteral != "" && tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}
