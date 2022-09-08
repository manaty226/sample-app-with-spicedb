package config

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConfig(t *testing.T) {
	t.Setenv("NAME_SPACE_FILE", "../../.config/namespace.txt")
	t.Setenv("AUTHN_FILE", "../../.config/user.csv")

	cfg, err := New()
	require.NoError(t, err)

	require.Equal(t, "127.0.0.1", cfg.SpiceHost)
	require.Equal(t, 3000, cfg.SpicePort)
}
