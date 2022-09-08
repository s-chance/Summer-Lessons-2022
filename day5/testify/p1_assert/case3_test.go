package p1_assert

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAssertCase3(t *testing.T) {
	assertion := assert.New(t)

	assertion.Contains("hello world", "world")
	assertion.NotContains("hello world", "hello")

	s := make([]int, 4)
	assertion.Len(s, 3)

	assertion.FileExists("/Users/patrick_star/GolandProjects/day5/testify/p1_assert/case3_test.go")
	assertion.DirExists("/Users/patrick_star/GolandProjects/day5/testify/p1_assert")
}
