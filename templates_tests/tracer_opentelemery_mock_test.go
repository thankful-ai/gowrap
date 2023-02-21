package templatestests

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

//go:generate minimock -i go.opentelemetry.io/otel/trace.Tracer -o ./tracer_opentelemery_mock_test.go -n OpentelemetryTracerMock

import (
	"context"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
	mm_trace "go.opentelemetry.io/otel/trace"
)

// OpentelemetryTracerMock implements trace.Tracer
type OpentelemetryTracerMock struct {
	t minimock.Tester

	funcStart          func(ctx context.Context, spanName string, opts ...mm_trace.SpanStartOption) (c2 context.Context, s1 mm_trace.Span)
	inspectFuncStart   func(ctx context.Context, spanName string, opts ...mm_trace.SpanStartOption)
	afterStartCounter  uint64
	beforeStartCounter uint64
	StartMock          mOpentelemetryTracerMockStart
}

// NewOpentelemetryTracerMock returns a mock for trace.Tracer
func NewOpentelemetryTracerMock(t minimock.Tester) *OpentelemetryTracerMock {
	m := &OpentelemetryTracerMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.StartMock = mOpentelemetryTracerMockStart{mock: m}
	m.StartMock.callArgs = []*OpentelemetryTracerMockStartParams{}

	return m
}

type mOpentelemetryTracerMockStart struct {
	mock               *OpentelemetryTracerMock
	defaultExpectation *OpentelemetryTracerMockStartExpectation
	expectations       []*OpentelemetryTracerMockStartExpectation

	callArgs []*OpentelemetryTracerMockStartParams
	mutex    sync.RWMutex
}

// OpentelemetryTracerMockStartExpectation specifies expectation struct of the Tracer.Start
type OpentelemetryTracerMockStartExpectation struct {
	mock    *OpentelemetryTracerMock
	params  *OpentelemetryTracerMockStartParams
	results *OpentelemetryTracerMockStartResults
	Counter uint64
}

// OpentelemetryTracerMockStartParams contains parameters of the Tracer.Start
type OpentelemetryTracerMockStartParams struct {
	ctx      context.Context
	spanName string
	opts     []mm_trace.SpanStartOption
}

// OpentelemetryTracerMockStartResults contains results of the Tracer.Start
type OpentelemetryTracerMockStartResults struct {
	c2 context.Context
	s1 mm_trace.Span
}

// Expect sets up expected params for Tracer.Start
func (mmStart *mOpentelemetryTracerMockStart) Expect(ctx context.Context, spanName string, opts ...mm_trace.SpanStartOption) *mOpentelemetryTracerMockStart {
	if mmStart.mock.funcStart != nil {
		mmStart.mock.t.Fatalf("OpentelemetryTracerMock.Start mock is already set by Set")
	}

	if mmStart.defaultExpectation == nil {
		mmStart.defaultExpectation = &OpentelemetryTracerMockStartExpectation{}
	}

	mmStart.defaultExpectation.params = &OpentelemetryTracerMockStartParams{ctx, spanName, opts}
	for _, e := range mmStart.expectations {
		if minimock.Equal(e.params, mmStart.defaultExpectation.params) {
			mmStart.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmStart.defaultExpectation.params)
		}
	}

	return mmStart
}

// Inspect accepts an inspector function that has same arguments as the Tracer.Start
func (mmStart *mOpentelemetryTracerMockStart) Inspect(f func(ctx context.Context, spanName string, opts ...mm_trace.SpanStartOption)) *mOpentelemetryTracerMockStart {
	if mmStart.mock.inspectFuncStart != nil {
		mmStart.mock.t.Fatalf("Inspect function is already set for OpentelemetryTracerMock.Start")
	}

	mmStart.mock.inspectFuncStart = f

	return mmStart
}

// Return sets up results that will be returned by Tracer.Start
func (mmStart *mOpentelemetryTracerMockStart) Return(c2 context.Context, s1 mm_trace.Span) *OpentelemetryTracerMock {
	if mmStart.mock.funcStart != nil {
		mmStart.mock.t.Fatalf("OpentelemetryTracerMock.Start mock is already set by Set")
	}

	if mmStart.defaultExpectation == nil {
		mmStart.defaultExpectation = &OpentelemetryTracerMockStartExpectation{mock: mmStart.mock}
	}
	mmStart.defaultExpectation.results = &OpentelemetryTracerMockStartResults{c2, s1}
	return mmStart.mock
}

// Set uses given function f to mock the Tracer.Start method
func (mmStart *mOpentelemetryTracerMockStart) Set(f func(ctx context.Context, spanName string, opts ...mm_trace.SpanStartOption) (c2 context.Context, s1 mm_trace.Span)) *OpentelemetryTracerMock {
	if mmStart.defaultExpectation != nil {
		mmStart.mock.t.Fatalf("Default expectation is already set for the Tracer.Start method")
	}

	if len(mmStart.expectations) > 0 {
		mmStart.mock.t.Fatalf("Some expectations are already set for the Tracer.Start method")
	}

	mmStart.mock.funcStart = f
	return mmStart.mock
}

// When sets expectation for the Tracer.Start which will trigger the result defined by the following
// Then helper
func (mmStart *mOpentelemetryTracerMockStart) When(ctx context.Context, spanName string, opts ...mm_trace.SpanStartOption) *OpentelemetryTracerMockStartExpectation {
	if mmStart.mock.funcStart != nil {
		mmStart.mock.t.Fatalf("OpentelemetryTracerMock.Start mock is already set by Set")
	}

	expectation := &OpentelemetryTracerMockStartExpectation{
		mock:   mmStart.mock,
		params: &OpentelemetryTracerMockStartParams{ctx, spanName, opts},
	}
	mmStart.expectations = append(mmStart.expectations, expectation)
	return expectation
}

// Then sets up Tracer.Start return parameters for the expectation previously defined by the When method
func (e *OpentelemetryTracerMockStartExpectation) Then(c2 context.Context, s1 mm_trace.Span) *OpentelemetryTracerMock {
	e.results = &OpentelemetryTracerMockStartResults{c2, s1}
	return e.mock
}

// Start implements trace.Tracer
func (mmStart *OpentelemetryTracerMock) Start(ctx context.Context, spanName string, opts ...mm_trace.SpanStartOption) (c2 context.Context, s1 mm_trace.Span) {
	mm_atomic.AddUint64(&mmStart.beforeStartCounter, 1)
	defer mm_atomic.AddUint64(&mmStart.afterStartCounter, 1)

	if mmStart.inspectFuncStart != nil {
		mmStart.inspectFuncStart(ctx, spanName, opts...)
	}

	mm_params := &OpentelemetryTracerMockStartParams{ctx, spanName, opts}

	// Record call args
	mmStart.StartMock.mutex.Lock()
	mmStart.StartMock.callArgs = append(mmStart.StartMock.callArgs, mm_params)
	mmStart.StartMock.mutex.Unlock()

	for _, e := range mmStart.StartMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.c2, e.results.s1
		}
	}

	if mmStart.StartMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmStart.StartMock.defaultExpectation.Counter, 1)
		mm_want := mmStart.StartMock.defaultExpectation.params
		mm_got := OpentelemetryTracerMockStartParams{ctx, spanName, opts}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmStart.t.Errorf("OpentelemetryTracerMock.Start got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmStart.StartMock.defaultExpectation.results
		if mm_results == nil {
			mmStart.t.Fatal("No results are set for the OpentelemetryTracerMock.Start")
		}
		return (*mm_results).c2, (*mm_results).s1
	}
	if mmStart.funcStart != nil {
		return mmStart.funcStart(ctx, spanName, opts...)
	}
	mmStart.t.Fatalf("Unexpected call to OpentelemetryTracerMock.Start. %v %v %v", ctx, spanName, opts)
	return
}

// StartAfterCounter returns a count of finished OpentelemetryTracerMock.Start invocations
func (mmStart *OpentelemetryTracerMock) StartAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmStart.afterStartCounter)
}

// StartBeforeCounter returns a count of OpentelemetryTracerMock.Start invocations
func (mmStart *OpentelemetryTracerMock) StartBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmStart.beforeStartCounter)
}

// Calls returns a list of arguments used in each call to OpentelemetryTracerMock.Start.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmStart *mOpentelemetryTracerMockStart) Calls() []*OpentelemetryTracerMockStartParams {
	mmStart.mutex.RLock()

	argCopy := make([]*OpentelemetryTracerMockStartParams, len(mmStart.callArgs))
	copy(argCopy, mmStart.callArgs)

	mmStart.mutex.RUnlock()

	return argCopy
}

// MinimockStartDone returns true if the count of the Start invocations corresponds
// the number of defined expectations
func (m *OpentelemetryTracerMock) MinimockStartDone() bool {
	for _, e := range m.StartMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.StartMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterStartCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcStart != nil && mm_atomic.LoadUint64(&m.afterStartCounter) < 1 {
		return false
	}
	return true
}

// MinimockStartInspect logs each unmet expectation
func (m *OpentelemetryTracerMock) MinimockStartInspect() {
	for _, e := range m.StartMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to OpentelemetryTracerMock.Start with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.StartMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterStartCounter) < 1 {
		if m.StartMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to OpentelemetryTracerMock.Start")
		} else {
			m.t.Errorf("Expected call to OpentelemetryTracerMock.Start with params: %#v", *m.StartMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcStart != nil && mm_atomic.LoadUint64(&m.afterStartCounter) < 1 {
		m.t.Error("Expected call to OpentelemetryTracerMock.Start")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *OpentelemetryTracerMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockStartInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *OpentelemetryTracerMock) MinimockWait(timeout mm_time.Duration) {
	timeoutCh := mm_time.After(timeout)
	for {
		if m.minimockDone() {
			return
		}
		select {
		case <-timeoutCh:
			m.MinimockFinish()
			return
		case <-mm_time.After(10 * mm_time.Millisecond):
		}
	}
}

func (m *OpentelemetryTracerMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockStartDone()
}
