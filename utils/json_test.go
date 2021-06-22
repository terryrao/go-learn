package utils

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJson(t *testing.T) {

	_, err := json.Marshal(true)

	assert.Nil(t, err)
}
