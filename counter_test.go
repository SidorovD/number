package counter_test

import (
	. "counter"
	"testing"
)

func assert(t *testing.T, want, got uint) {
	if want != got {
		t.Errorf("want: %v, got: %v", want, got)
	}
}

func TestNew(t *testing.T) {
	c := New()
	assert(t, c.Val(), 0)
	assert(t, c.MaxVal(), MaxUint)
}

func TestIncrement(t *testing.T) {
	c := New()
	c = c.Inc()
	assert(t, c.Val(), 1)
}

func TestSetMaxValue(t *testing.T) {
	c := New()
	c = c.SetMaxValue(2)
	assert(t, c.MaxVal(), 2)
}

func TestIncrementOverflow(t *testing.T) {
	c := New()
	c = c.SetMaxValue(1)
	c = c.Inc()

	// overflow number
	c = c.Inc()
	assert(t, c.Val(), 0)
}

func TestSetMaxValue_IfNewMaxLessThanValueThenSetValueToZero(t *testing.T) {
	c := New()
	c = c.Inc()
	c = c.Inc()
	c = c.Inc()

	c = c.SetMaxValue(2)
	assert(t, c.Val(), 0)
}
