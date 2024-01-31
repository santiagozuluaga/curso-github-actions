package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHandler(t *testing.T) {
	c := require.New(t)

	backupVar := os.Getenv("GITHUB_ACTOR")

	defer func() {
		os.Setenv("GITHUB_ACTOR", backupVar)
	}()

	os.Clearenv()

	c.Equal("", getEnvString("GITHUB_ACTOR", ""))
	c.Equal("santiagozuluaga", getEnvString("GITHUB_ACTOR", "santiagozuluaga"))

	os.Setenv("GITHUB_ACTOR", "szuluaga")
	c.Equal("szuluaga", getEnvString("GITHUB_ACTOR", ""))
}
