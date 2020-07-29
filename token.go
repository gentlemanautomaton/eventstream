package eventstream

// TokenType indicates the type of an event stream token returned by scanner.
type TokenType int

// Token types.
const (
	EndToken     TokenType = 0
	CommentToken TokenType = 1
	FieldToken   TokenType = 2
)

// Token holds a reference to buffered line data read by Scanner. It is safe
// to pass tokens by value.
//
// A token is only valid until the next call to scanner.Scan, after which its
// underlying data may be overwritten.
type Token struct {
	line  []byte
	colon int
}

// Type returns the type of the token.
func (t Token) Type() TokenType {
	switch {
	case len(t.line) == 0:
		return EndToken
	case t.colon == 0:
		return CommentToken
	default:
		return FieldToken
	}
}

// Comment returns the token as a comment.
func (t Token) Comment() Comment {
	// Check for non-comment and empty comment
	if t.colon != 0 || len(t.line) <= 1 {
		return nil
	}
	return Comment(t.line[1:])
}

// Field interprets the token as a field and returns it.
func (t Token) Field() Field {
	switch {
	case t.colon == 0:
		return nil // Comment
	case t.colon < 0:
		return Field(t.line) // Field without value
	default:
		return Field(t.line[0:t.colon]) // Field with value
	}
}

// Value interprets the token as a field and returns its value, if present.
func (t Token) Value() Value {
	// Check for comment (colon == 0) and field without value (colon == -1)
	if t.colon <= 0 {
		return nil
	}

	length := len(t.line)

	// Don't include the colon delimeter
	start := t.colon + 1

	// Check for empty value (without space prefix)
	if start >= length {
		return nil
	}

	// Don't include the space prefix if present
	if t.line[start] == ' ' {
		start++
	}

	// Check for empty value (with space prefix)
	if start >= length {
		return nil
	}

	return Value(t.line[start:])
}
