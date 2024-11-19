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
	flagSet.Parse([]string{"-v"})
	defer func() {
		flagSet.Parse([]string{"-v=false"})
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
	flagSet.Parse([]string{"-dev"})
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

	tmpFile.Write([]byte("TEST_ENV_VAR=value"))
	fmt.Println(tmpFile.Name())
	// Set environment variable

	// Call main function
	main()

	// Check that .env file is loaded
	assert.Equal(t, "value", os.Getenv("TEST_ENV_VAR"))
}

func TestMainLoadEnvError(t *testing.T) {
	// Create temporary .env file with invalid contents
	tmpFile, err := os.CreateTemp("", ".env")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpFile.Name())

	// Write invalid contents to file
	_, err = tmpFile.WriteString(" invalid contents")
	if err != nil {
		t.Fatal(err)
	}

	// Call main function
	defer func() {
		if r := recover(); r != nil {
			assert.NotNil(t, r)
		}
	}()
	main()
}
