package blocks_test

import (
	"testing"

	"github.com/rentziass/gh-runmd/blocks"
	"github.com/stretchr/testify/assert"
)

func TestForPaths(t *testing.T) {
	linux := &blocks.Block{Paths: []string{"linux"}}
	macos := &blocks.Block{Paths: []string{"macos"}}
	windows := &blocks.Block{Paths: []string{"windows"}}

	allBlocks := blocks.Blocks{linux, macos, windows}

	tests := []struct {
		params []string
		want   blocks.Blocks
	}{
		{
			params: []string{"linux"},
			want:   blocks.Blocks{linux},
		},
		{
			params: []string{"other"},
			want:   nil,
		},
		{
			params: []string{"linux", "macos", "other"},
			want:   blocks.Blocks{linux, macos},
		},
		{
			params: []string{"other", "linux"},
			want:   blocks.Blocks{linux},
		},
	}

	for _, tt := range tests {
		result := allBlocks.ForPaths(tt.params...)
		assert.Equal(t, tt.want, result)
	}
}
