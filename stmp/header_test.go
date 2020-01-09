package stmp

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"strings"
	"testing"
)

func TestNewHeader(t *testing.T) {
	h := NewHeader()
	h.Set("Escape: A:B\nC;%", "B%C\n:D")
	h.Add("Two", "One")
	h.Add("Two", "Two")
	h.Set("Normal", "Value")
	expects := []string{
		"Escape%3A A%3AB%0AC%3B%25:B%25C%0A%3AD",
		"Two:One",
		"Two:Two",
		"Normal:Value",
	}
	sort.Strings(expects)
	actuals := strings.Split(string(h.Marshal()), ";")
	sort.Strings(actuals)
	assert.Equal(t, expects, actuals, "marshal")
	h1 := NewHeader()
	err := h1.Unmarshal(h.Marshal())
	assert.Nil(t, err, "unmarshal without error")
	assert.Equal(t, h, h1, "unmarshal")
}
