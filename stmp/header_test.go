// Copyright 2020 acrazing <joking.young@gmail.com>. All rights reserved.
// Since 2020-01-05 22:00:06
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
	assert.Equal(t, "escaped%3A A%3AB%0Ac:B%25C%0A:D\ntwo:One\ntwo:Two\nnormal:Value", string(h.Marshal()), "header marshal")
}
