package stmp

import (
	"math/rand"
	"time"
)

type Backoff interface {
	Next() (time.Duration, bool)
	Reset()
}

type linearBackoff struct {
	base   time.Duration
	factor time.Duration
	jitter int
	limit  int
	count  int
}

func (l *linearBackoff) Next() (time.Duration, bool) {
	if l.count >= l.limit {
		return 0, false
	}
	l.count += 1
	return (l.base + l.factor*time.Duration(l.count)) * time.Duration(100+l.jitter-rand.Intn(l.jitter<<1)) / 100, true
}

func (l *linearBackoff) Reset() {
	l.count = 0
}

// jitter is the random time to add or del range limit, from 0 to 100
func NewBackoff(base, factor time.Duration, jitter int, limit int) Backoff {
	return &linearBackoff{
		base:   base,
		factor: factor,
		jitter: jitter,
		limit:  limit,
	}
}
