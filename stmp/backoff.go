package stmp

import (
	"math/rand"
	"time"
)

type Backoff interface {
	// the next wait time
	// if should stop, the second value should be false
	// else it should be true
	Next() (wait time.Duration, count int, ok bool)
	// reset the count to 0
	Reset()
}

type linearBackoff struct {
	base   time.Duration
	factor time.Duration
	jitter int
	limit  int
	count  int
}

func (l *linearBackoff) Next() (time.Duration, int, bool) {
	if l.count >= l.limit {
		return 0, 0, false
	}
	l.count += 1
	return (l.base + l.factor*time.Duration(l.count)) * time.Duration(100+l.jitter-rand.Intn(l.jitter<<1)) / 100, l.count, true
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
