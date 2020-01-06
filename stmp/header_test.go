package stmp

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewHeader(t *testing.T) {
	h := NewHeader()
	h.Set("Escape: A:B\nC", "B%C\n:D")
	h.Set("Two", "One", "Two")
	h.Set("Normal", "Value")
	assert.Equal(t, "escape%3A a%3Ab%0Ac:B%25C%0A:D\ntwo:One\ntwo:Two\nnormal:Value", string(h.Marshal()), "marshal")
	h1 := NewHeader()
	err := h1.Unmarshal(h.Marshal())
	assert.Nil(t, err, "unmarshal without error")
	assert.Equal(t, h, h1, "unmarshal")
}
