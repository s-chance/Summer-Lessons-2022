package p1_assert

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAssertCase2(t *testing.T) {
	assertion := assert.New(t)

	name := "Bob"
	age := 10

	assertion.Equalf("bob", name, "they should be equal")
	assertion.NotEqualf(10, age, "they should not be equal")

	assertion.Nil("", "they should be nil")
	assertion.NotNil(nil, "they should not be nil")

	//var i interface{}
	assertion.Empty("i", "they should be empty") //可以是0, "", false, 0.0, 不能是" "和' '
	assertion.NotEmpty("", "they should not be empty")

	err := errors.New("new error")
	var err2 error
	assertion.Error(err, "they should be a error")
	assertion.NoError(err2, "they should not be a error")

	var flag bool
	assertion.True(flag, "they should be true")
	assertion.False(flag, "they should be false")
}
