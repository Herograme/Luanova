package luanova

import "unicode"

type Token struct {
	Type    int
	Literal string
}

type Lexer struct {
	input   string
	pos     int
	readPos int
	ch      byte
}

func NewLexer(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPos >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPos]
	}
	l.pos = l.readPos
	l.readPos++
}

func (l *Lexer) peekChar() byte {
	if l.readPos >= len(l.input) {
		return 0
	}
	return l.input[l.readPos]
}

func (l *Lexer) NextToken() Token {
	var tok Token

	l.skipWhitespace()

	switch l.ch {
	case '+':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = Token{Type: PlusAssign, Literal: string(ch) + string(l.ch)}
		} else {
			tok = Token{Type: Plus, Literal: string(l.ch)}
		}
	case '-':
		switch l.peekChar() {
		case '=':
			ch := l.ch
			l.readChar()
			tok = Token{Type: SubAssign, Literal: string(ch) + string(l.ch)}
		case '>':
			ch := l.ch
			l.readChar()
			tok = Token{Type: Arrow, Literal: string(ch) + string(l.ch)}
		case '-':
			tok.Type = Comment
			tok.Literal = l.readLineComment()
			return tok
		case '*':
			tok.Type = CommentBlock
			tok.Literal = l.readBlockComment()
			return tok
		default:
			tok = Token{Type: Sub, Literal: string(l.ch)}
		}
	case '*':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = Token{Type: MultiAssign, Literal: string(ch) + string(l.ch)}
		} else {
			tok = Token{Type: Multi, Literal: string(l.ch)}
		}
	case '/':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = Token{Type: DivAssign, Literal: string(ch) + string(l.ch)}
		} else {
			tok = Token{Type: Div, Literal: string(l.ch)}
		}
	case '%':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = Token{Type: ModAssign, Literal: string(ch) + string(l.ch)}
		} else {
			tok = Token{Type: Mod, Literal: string(l.ch)}
		}
	case '^':
		tok = Token{Type: Po, Literal: string(l.ch)}
	case '.':
		if l.peekChar() == '.' {
			l.readChar()
			if l.peekChar() == '.' {
				l.readChar()
				tok = Token{Type: Dots, Literal: "..."}
			} else {
				tok = Token{Type: Concat, Literal: ".."}
			}
		} else {
			tok = Token{Type: Dot, Literal: string(l.ch)}
		}
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = Token{Type: Equal, Literal: string(ch) + string(l.ch)}
		} else {
			tok = Token{Type: Assign, Literal: string(l.ch)}
		}

	case '~':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = Token{Type: NotEqual, Literal: string(ch) + string(l.ch)}
		} else {
			tok = Token{Type: Illegal, Literal: string(l.ch)}
		}
	case '<':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = Token{Type: LessEqual, Literal: string(ch) + string(l.ch)}
		} else {
			tok = Token{Type: Less, Literal: string(l.ch)}
		}
	case '>':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = Token{Type: GreaterEqual, Literal: string(ch) + string(l.ch)}
		} else {
			tok = Token{Type: Greater, Literal: string(l.ch)}
		}
	case '(':
		tok = Token{Type: LParen, Literal: string(l.ch)}
	case ')':
		tok = Token{Type: RParen, Literal: string(l.ch)}
	case '{':
		tok = Token{Type: LBrace, Literal: string(l.ch)}
	case '}':
		tok = Token{Type: RBrace, Literal: string(l.ch)}
	case '[':
		tok = Token{Type: LBrack, Literal: string(l.ch)}
	case ']':
		tok = Token{Type: RBrack, Literal: string(l.ch)}
	case ',':
		tok = Token{Type: Comma, Literal: string(l.ch)}
	case ':':
		tok = Token{Type: Colom, Literal: string(l.ch)}
	case ';':
		tok = Token{Type: Semi, Literal: string(l.ch)}
	case '"':
		tok = Token{Type: StringDelim, Literal: l.readString()}
	case '?':
		tok = Token{Type: Question, Literal: string(l.ch)}
	case 0:
		tok.Literal = ""
		tok.Type = EOF
	default:
		if isLetter(l.ch) {
			literal := l.readIdentifier()
			tok.Type = LookupIdent(literal)
			tok.Literal = literal
			return tok
		} else if isDigit(l.ch) {
			literal := l.readNumber()
			tok.Literal = literal
			tok.Type = Literal
			return tok
		} else {
			tok = Token{Type: Illegal, Literal: string(l.ch)}
		}
	}

	l.readChar()
	return tok
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) readIdentifier() string {
	start := l.pos
	for isLetter(l.ch) || isDigit(l.ch) {
		l.readChar()
	}
	return l.input[start:l.pos]
}
func (l *Lexer) readLineComment() string {
	start := l.pos
	for l.ch != '\n' && l.ch != 0 {
		l.readChar()
	}
	return l.input[start:l.pos]
}

func isLetter(ch byte) bool {
	return unicode.IsLetter(rune(ch)) || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func LookupIdent(ident string) int {
	var token int
	switch ident {
	case "function":
		token = Function
	case "local":
		token = Local
	case "if":
		token = If
	case "else":
		token = Else
	case "elseif":
		token = ElseIf
	case "and":
		token = And
	case "or":
		token = Or
	case "not":
		token = Not
	case "while":
		token = While
	case "for":
		token = For
	case "return":
		token = Return
	case "break":
		token = Break
	case "continue":
		token = Continue
	case "true":
		token = True
	case "false":
		token = False
	case "do":
		token = Do
	case "in":
		token = In
	case "then":
		token = Then
	case "end":
		token = End

	default:
		token = Literal
	}

	return token
}

func (l *Lexer) readString() string {
	l.readChar()
	start := l.pos
	for l.ch != '"' && l.ch != 0 {
		l.readChar()
	}
	return l.input[start:l.pos]
}

func (l *Lexer) readBlockComment() string {
	var start int
	if l.pos > 0 {
		start = l.pos - 1
	} else {
		start = 0
	}

	l.readChar()

	for {
		if l.ch == '*' && l.peekChar() == '-' {
			l.readChar()
			l.readChar()
			break
		}
		if l.ch == 0 {
			break
		}
		l.readChar()
	}

	return l.input[start:l.pos]
}

func (l *Lexer) readNumber() string {
	start := l.pos
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[start:l.pos]
}
