package blocks

import (
	"golang.org/x/exp/slices"
)

type Blocks []*Block

type Block struct {
	Content string
	Paths   []string
}

func (blocks Blocks) ForPaths(paths ...string) Blocks {
	var result Blocks
	for _, block := range blocks {
		if slices.ContainsFunc(paths, func(path string) bool {
			return slices.Contains(block.Paths, path)
		}) {
			result = append(result, block)
		}
	}

	return result
}

func (blocks Blocks) ToScript() string {
	var result string
	for _, block := range blocks {
		result += block.Content
	}

	return result
}
