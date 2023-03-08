package wiregen_test

import (
	"fmt"
	"testing"
	"wire-example/internal/config"
	wiregen "wire-example/internal/wire-gen"

	"github.com/stretchr/testify/require"
)

func TestProvideServer(t *testing.T) {
	cfg := config.Config{ImportantConfigValue: "all good"}

	server, err := wiregen.ProvideServer(cfg)
	require.NoError(t, err)
	fmt.Println(server.Azure.GetClientName())
	require.NoError(t, server.Azure.CheckDBConnection())

	serverMock, err := wiregen.ProvideServerMock(cfg)
	require.NoError(t, err)
	fmt.Println(serverMock.Azure.GetClientName())
	require.NoError(t, serverMock.Azure.CheckDBConnection())
}
