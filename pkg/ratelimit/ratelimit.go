package ratelimit

import "golang.org/x/time/rate"

type Limiter struct {
	Burst      int
	LimitAllow int
	Limiter    *rate.Limiter
}

func NewLimitter(burst, limit int) *Limiter {
	limiter := rate.NewLimiter(rate.Limit(limit), burst)
	return &Limiter{
		Limiter: limiter,
	}
}

func (l *Limiter) Limit() bool {
	return l.Limiter.Allow()
}
