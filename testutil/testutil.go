package testutil

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func LoadFile(t *testing.T, path string) []byte {
	t.Helper()

	b, err := os.ReadFile(path)
	require.NoError(t, err)
	return b
}
