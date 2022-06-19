package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testTable struct {
	name string
	params string
	returns string
}

func TestSample(t *testing.T) {
	result := Sample("Angga")

	if result != "Hello Angga" {
		t.Fatal("return harusnya Hello Angga")
	}
}

func TestSamples(t *testing.T) {
	result := Sample("Angga")
	assert.Equal(t, "Hello Angga", result, "salah")
}

func TestSampleSub(t *testing.T) {
	t.Run("params angga", func(t *testing.T) {
		result := Sample("Angga")
		assert.Equal(t, "Hello Angga", result, "salah")
	})
}

func TestTable (t *testing.	T) {
	var testTb = []testTable{
		{
			name: "angga",
			params: "angga",
			returns: "hello angga",
		},
		{
			name: "devri",
			params: "devri",
			returns: "hello devri",
		},
	}

	for _, val := range testTb {
		t.Run(val.name, func(t *testing.T) {
			result := Sample(val.params)
			assert.Equal(t, val.returns, result, "salah")
		})
	}
}