package blocks

import (
	"fmt"
	"io"

	"gitlab.com/golang-commonmark/markdown"
)

func From(r io.Reader) (Blocks, error) {
	m := markdown.New()

	b, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}

	fmt.Println("b:", string(b))

	tokens := m.Parse(b)
	fmt.Println("tokens:", tokens)
	for _, token := range tokens {
		fmt.Printf("%T %+v\n", token, token)
	}

	// ...
	return nil, nil
}
