package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_version(t *testing.T) {
	flagTrue := true
	v = &flagTrue
	main()
	assert.True(t, *v)
}

func Test_dev(t *testing.T) {
	flagTrue := true
	isDev = &flagTrue
	main()
	assert.True(t, *isDev)
}
