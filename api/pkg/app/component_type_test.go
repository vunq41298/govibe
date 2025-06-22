package app

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestComponentType_String(t *testing.T) {
	require.Equal(t, "api", ComponentTypeAPI.String())
	require.Equal(t, "job", ComponentTypeJob.String())
	require.Equal(t, "consumer", ComponentTypeConsumer.String())
	require.Equal(t, "abcd", ComponentType("abcd").String())
	require.Equal(t, "", ComponentType("").String())
}

func TestComponentType_Valid(t *testing.T) {
	require.True(t, ComponentTypeAPI.Valid())
	require.True(t, ComponentTypeJob.Valid())
	require.True(t, ComponentTypeConsumer.Valid())
	require.False(t, ComponentType("abcd").Valid())
	require.False(t, ComponentType("").Valid())
}
