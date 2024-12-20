package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MainVersionFlag(t *testing.T) {
	// Set version flag
	flagSet := flag.NewFlagSet("test", flag.ExitOnError)
	flagSet.BoolVar(v, "v", true, "version")
	err := flagSet.Parse([]string{"-v"})
	assert.Nil(t, err)
	defer func() {
		err = flagSet.Parse([]string{"-v=false"})
		assert.Nil(t, err)
		assert.False(t, *v)
	}()
	// Capture output
	// keep backup of the real stdout
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	w.Close()
	out, _ := io.ReadAll(r)
	os.Stdout = rescueStdout

	fmt.Printf("Captured: %s", out)

	// Check output
	assert.Contains(t, string(out), Version)
}

func Test_MainDevFlag(t *testing.T) {
	// Set dev flag
	flagSet := flag.NewFlagSet("test", flag.ExitOnError)
	flagSet.BoolVar(isDev, "dev", false, "dev mode")
	err := flagSet.Parse([]string{"-dev"})
	assert.Nil(t, err)
	// Call main function
	main()

	// Check that dev flag is set
	assert.True(t, *isDev)
}

func Test_MainEnvFileExists(t *testing.T) {
	// Create temporary .env file
	path, err := os.Getwd()
	assert.Nil(t, err)
	envFilePath := path + "/.env"
	tmpFile, err := os.Create(envFilePath)
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(envFilePath)

	_, err = tmpFile.Write([]byte("TEST_ENV_VAR=value"))
	assert.Nil(t, err)
	fmt.Println(tmpFile.Name())
	// Set environment variable

	// Call main function
	main()

	// Check that .env file is loaded
	assert.Equal(t, "value", os.Getenv("TEST_ENV_VAR"))
}

func Test_handlerErr(t *testing.T) {
	assert.Panics(t, func() {
		err := fmt.Errorf("test error")
		handler(err)
	}, "test error")
}
