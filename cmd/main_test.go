package main

import (
	"flag"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestName(t *testing.T) {
	appName = "admin"
	assert.Equal(t, "admin", GetAppName())
}

func TestParseFlagsWrong(t *testing.T) {
	appName = ""
	err := parseFlags()
	assert.EqualErrorf(t, err, ErrNoFlags, "Error shold be $v, got: %v", ErrNoFlags, err)
}

func TestParseFlagsRight(t *testing.T) {
	flag.StringVar(&appName, "app", "", "application name [admin|redirect]")
	os.Args = append(os.Args, "-app", "admin")
	err := parseFlags()
	require.NoError(t, err)
}

func TestStart(t *testing.T) {
	appName = "admin"
	os.Chdir("../")
	t.Run("Test right logs dir", func(t *testing.T) {
		appName = "admin"
		_, f, err := start()
		defer f.Close()
		require.NoError(t, err)
	})
	oldPath := logPath
	t.Run("Test wrong logs dir", func(t *testing.T) {
		logPath = "/root/"
		_, f, err := start()
		defer f.Close()
		require.Error(t, err)
	})
	t.Run("Test no such app", func(t *testing.T) {
		logPath = oldPath
		appName = "NoSuch"
		_, f, err := start()
		defer f.Close()
		require.Error(t, err)
	})
}
