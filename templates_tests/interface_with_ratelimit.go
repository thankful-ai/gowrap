package templatestests

// DO NOT EDIT!
// This code is generated with http://github.com/hexdigest/gowrap tool
// using ../templates/ratelimit template

//go:generate gowrap gen -p github.com/hexdigest/gowrap/templates_tests -i TestInterface -t ../templates/ratelimit -o interface_with_ratelimit.go

import (
	"context"
	"time"
)

// TestInterfaceWithRateLimit implements TestInterface
type TestInterfaceWithRateLimit struct {
	_base  TestInterface
	_ticks chan time.Time
}

// NewTestInterfaceWithRateLimit instruments an implementation of the TestInterface with rate limiting
func NewTestInterfaceWithRateLimit(base TestInterface, burst int, rps float64) *TestInterfaceWithRateLimit {
	d := &TestInterfaceWithRateLimit{
		_base:  base,
		_ticks: make(chan time.Time, burst),
	}

	now := time.Now()
	for i := 0; i < burst; i++ {
		d._ticks <- now
	}

	delay := time.Duration(float64(time.Second) / rps)

	go func() {
		for t := range time.Tick(delay) {
			d._ticks <- t
		}
	}()

	return d
}

// Channels implements TestInterface
func (_d *TestInterfaceWithRateLimit) Channels(chA chan bool, chB chan<- bool, chanC <-chan bool) {
	<-_d._ticks

	_d._base.Channels(chA, chB, chanC)
	return
}

// F implements TestInterface
func (_d *TestInterfaceWithRateLimit) F(ctx context.Context, a1 string, a2 ...string) (result1 string, result2 string, err error) {
	select {
	case <-ctx.Done():
		err = ctx.Err()
		return
	case <-_d._ticks:
	}

	return _d._base.F(ctx, a1, a2...)
}

// NoError implements TestInterface
func (_d *TestInterfaceWithRateLimit) NoError(s1 string) (s2 string) {
	<-_d._ticks

	return _d._base.NoError(s1)
}

// NoParamsOrResults implements TestInterface
func (_d *TestInterfaceWithRateLimit) NoParamsOrResults() {
	<-_d._ticks

	_d._base.NoParamsOrResults()
	return
}
