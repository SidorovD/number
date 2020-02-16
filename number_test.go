package number_test

import (
	. "number"
	"testing"
)

func assert(t *testing.T, want, got uint) {
	if want != got {
		t.Errorf("want: %v, got: %v", want, got)
	}
}

func TestNew(t *testing.T) {
	inc := New()
	assert(t, inc.Val(), 0)
	assert(t, inc.MaxVal(), Max)
}

func TestIncrement(t *testing.T) {
	inc := New()
	inc = inc.Inc()
	assert(t, inc.Val(), 1)
}

func TestSetMaxValue(t *testing.T) {
	num := New()
	num = num.SetMaxValue(2)
	assert(t, num.MaxVal(), 2)
}

func TestIncrementOverflow(t *testing.T) {
	inc := New()
	inc = inc.SetMaxValue(1)
	inc = inc.Inc()

	// overflow number
	inc = inc.Inc()
	assert(t, inc.Val(), 0)
}

func TestSetMaxValue_IfNewMaxLessThanValueThenSetValueToZero(t *testing.T) {
	inc := New()
	inc = inc.Inc()
	inc = inc.Inc()
	inc = inc.Inc()

	inc = inc.SetMaxValue(2)
	assert(t, inc.Val(), 0)
}

