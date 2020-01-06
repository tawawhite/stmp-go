package stmp_test

import (
	"github.com/acrazing/stmp-go/stmp"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHandshake(t *testing.T) {
	h := stmp.NewServerHandshake(0, nil, "")
	assert.Equal(t, []byte{'S', 'T', 'M', 'P', 0}, h.MarshalBinary(), "they should be equal")
}
