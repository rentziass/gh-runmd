package blocks

type Blocks []*Block

type Block struct {
	Content string
	Paths   []string
}

func (blocks Blocks) ForPaths(paths ...string) (Blocks, error) {
	// ...
	return nil, nil
}
