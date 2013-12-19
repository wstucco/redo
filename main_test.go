package main

import (
	"github.com/bmizerany/assert"
	"testing"
)

func TestReplaceBaseName(t *testing.T) {
	assert.Equal(t, replaceBaseName("redo.do", "default"), "default.do")
	assert.Equal(t, replaceBaseName("./redo.do", "default"), "default.do")
	assert.NotEqual(t, replaceBaseName("redo.do", "default"), ".default.do")
}

func TestMap(t *testing.T) {
	slice := StringSlice{"a", "b", "ab"}
	expected_result := StringSlice{"xa", "xb", "xab"}

	result := slice.Map(func(s string) string {
		return ("x" + s)
	})

	assert.Equal(t, result, expected_result)
}
