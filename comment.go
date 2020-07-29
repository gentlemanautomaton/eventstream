package eventstream

// Comment contains the UTF-8 encoded bytes of a comment.
type Comment []byte

// String returns a string representation of c.
func (c Comment) String() string {
	return string(c)
}
