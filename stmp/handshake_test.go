// Copyright 2020 acrazing <joking.young@gmail.com>. All rights reserved.
// Since 2020-01-05 21:56:24
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
