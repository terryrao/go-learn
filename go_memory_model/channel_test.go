package go_memory_model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBufferredChannel1(t *testing.T) {
	helloInGogroutine()
	assert.Equal(t, a, "hello world")
}
