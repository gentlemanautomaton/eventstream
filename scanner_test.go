package eventstream_test

import (
	"fmt"
	"strings"

	"github.com/gentlemanautomaton/eventstream"
)

func ExampleScanner() {
	data := `: test stream

data: first event
id: 1

data:second event
id

data:  third event`

	scanner := eventstream.NewScanner(strings.NewReader(data))
	for scanner.Scan() {
		t := scanner.Token()
		switch t.Type() {
		case eventstream.CommentToken:
			fmt.Printf("COMMENT: %s\n", t.Comment())
		case eventstream.FieldToken:
			fmt.Printf("FIELD: %s\n", t.Field())
			if v := t.Value(); v != nil {
				fmt.Printf("VALUE: %s\n", t.Value())
			}
		case eventstream.EndToken:
			fmt.Println()
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	// Output:
	// COMMENT:  test stream
	//
	// FIELD: data
	// VALUE: first event
	// FIELD: id
	// VALUE: 1
	//
	// FIELD: data
	// VALUE: second event
	// FIELD: id
	//
	// FIELD: data
	// VALUE:  third event
}
