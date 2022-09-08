package p1_assert

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAssertCase1(t *testing.T) {
	name := "Bob"
	age := 10

	assert.Equal(t, "bob", name, "they should be equal")
	assert.NotEqual(t, 10, age, "they should not be equal")
}
