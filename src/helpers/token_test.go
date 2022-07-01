package helpers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToken(t *testing.T) {
	username := "tokenuser"
	result := NewToken(username)
	assert.Equal(t, result, "expected create token")
}
 
