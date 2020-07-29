package eventstream

// Field contains the UTF-8 encoded bytes of an event field name.
type Field []byte

// Field returns a string representation of f.
func (f Field) String() string {
	return string(f)
}
