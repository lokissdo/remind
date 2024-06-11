package util

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	config, err := LoadConfig("../")
	require.NoError(t, err)

	require.Equal(t, config.DBSource, "postgres://remind_user:haFaMrzhWE0sFRMgHqdRrd6V4lBZniLY@dpg-cpduaf7sc6pc73970qk0-a.oregon-postgres.render.com/remind")
	require.Equal(t, config.ServerAddr.HTTP, "0.0.0.0:8080")
	require.Equal(t, config.ServerAddr.GRPC, "0.0.0.0:9090")
	require.Equal(t, config.PasetoConfig.TokenSymmetricKey, "12345678901234567890123456789012")
}
