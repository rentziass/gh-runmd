package blocks_test

import (
	"io"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/rentziass/gh-runmd/blocks"
)

func TestFrom(t *testing.T) {
	tests := []struct {
		name    string
		r       io.Reader
		want    blocks.Blocks
		wantErr bool
	}{
		{
			name: "empty",
			r:    strings.NewReader(""),
			want: blocks.Blocks{},
		},
		{
			name: "no blocks",
			r:    fileReader(t, "testdata/no_blocks.md"),
			want: blocks.Blocks{},
		},
		{
			name: "one block",
			r:    fileReader(t, "testdata/one_block.md"),
			want: blocks.Blocks{
				&blocks.Block{Content: "hello"},
			},
		},
		{
			name: "two blocks",
			r:    fileReader(t, "testdata/two_blocks.md"),
			want: blocks.Blocks{
				&blocks.Block{Content: "hello"},
				&blocks.Block{Content: "world"},
			},
		},
		{
			name: "two blocks with paths",
			r:    fileReader(t, "testdata/blocks_with_paths.md"),
			want: blocks.Blocks{
				&blocks.Block{
					Content: "hello",
					Paths:   []string{"macos", "linux"},
				},
				&blocks.Block{
					Content: "world",
					Paths:   []string{"windows"},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, err := blocks.From(tt.r)

			if !tt.wantErr {
				require.NoError(t, err)
			}

			if tt.wantErr {
				assert.Error(t, err)
			}

			assert.Equal(t, tt.want, b)
		})
	}
}

func fileReader(t *testing.T, path string) io.Reader {
	t.Helper()

	r, err := os.Open(path)
	require.NoError(t, err)

	return r
}
