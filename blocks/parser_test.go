package blocks_test

import (
	"fmt"
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
			want: nil,
		},
		{
			name: "no blocks",
			r:    fileReader(t, "testdata/no_blocks.md"),
			want: nil,
		},
		{
			name: "fenced block",
			r:    fileReader(t, "testdata/fenced_block.md"),
			want: blocks.Blocks{
				&blocks.Block{Content: "hello\nworld\n"},
			},
		},
		{
			name: "indented block",
			r:    fileReader(t, "testdata/indented_block.md"),
			want: nil,
		},
		{
			name: "two blocks",
			r:    fileReader(t, "testdata/two_blocks.md"),
			want: blocks.Blocks{
				&blocks.Block{Content: "hello\n"},
				&blocks.Block{Content: "world\n"},
			},
		},
		{
			name: "two blocks with paths",
			r:    fileReader(t, "testdata/blocks_with_paths.md"),
			want: blocks.Blocks{
				&blocks.Block{
					Content: "hello\n",
					Paths:   []string{"macos", "linux"},
				},
				&blocks.Block{
					Content: "world\n",
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
	fmt.Println("r:", r)
	require.NoError(t, err)

	t.Cleanup(func() {
		require.NoError(t, r.Close())
	})

	return r
}
