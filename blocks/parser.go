package blocks

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"gitlab.com/golang-commonmark/markdown"
)

func From(r io.Reader) (Blocks, error) {
	m := markdown.New()

	b, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}

	var blocks Blocks

	tokens := m.Parse(b)
	fmt.Println("tokens:", tokens)
	for _, token := range tokens {
		switch token := token.(type) {
		case *markdown.Fence:
			block, err := blockFromFence(token)
			if err != nil {
				return nil, err
			}

			if block == nil {
				continue
			}

			blocks = append(blocks, block)
		}
		fmt.Printf("%T %+v\n", token, token)
	}

	return blocks, nil
}

func blockFromFence(f *markdown.Fence) (*Block, error) {
	params, err := paramsFromFence(f)
	if err != nil {
		return nil, err
	}

	if params.Skip {
		return nil, nil
	}

	return &Block{
		Content: f.Content,
		Paths:   params.Paths,
	}, nil
}

type blockParams struct {
	Skip  bool     `json:"skip"`
	Paths []string `json:"paths"`
}

func paramsFromFence(f *markdown.Fence) (*blockParams, error) {
	jsonStarted := false
	paramsString := strings.TrimLeftFunc(f.Params, func(r rune) bool {
		if jsonStarted {
			return false
		}

		if r == '{' {
			jsonStarted = true
			return false
		}

		return true
	})

	params := &blockParams{}

	if len(paramsString) == 0 {
		return params, nil
	}

	err := json.Unmarshal([]byte(paramsString), params)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to unmarshal block params '%s': %w",
			paramsString,
			err,
		)
	}

	return params, nil
}
