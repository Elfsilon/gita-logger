package gita

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseLevel(t *testing.T) {
	lvl, err := ParseLevel("wArning")
	assert.Nil(t, err)
	require.Equal(t, lvl, WarningLevel)
}

func TestParseLevelUndefined(t *testing.T) {
	lvl, err := ParseLevel("dAnger")
	assert.NotNil(t, err)
	require.Equal(t, lvl, InfoLevel)
}

func TestParseLevelEmpty(t *testing.T) {
	lvl, err := ParseLevel("")
	assert.NotNil(t, err)
	require.Equal(t, lvl, InfoLevel)
}
