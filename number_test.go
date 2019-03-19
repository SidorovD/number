package number_test

import (
	. "number"
	"testing"
)

func TestIncrement(t *testing.T) {
	inc := New()

	if inc.Val() != 0 {
		t.Errorf("want: 0, got: %v", inc.Val())
	}

	inc = inc.Inc()

	if inc.Val() != 1 {
		t.Errorf("want: 1, got: %v", inc.Val())
	}
}

func TestIncrementOverflow(t *testing.T) {
	inc := New()

	for i := 0; i < Max; i++ {
		inc = inc.Inc()
	}

	// overflow number
	inc = inc.Inc()

	if v := inc.Val(); v != 0 {
		t.Errorf("want: 0, got: %v", v)
	}
}

func TestSetMaxValue(t *testing.T) {
	zero := New()
	zero = zero.SetMaxValue(2)

	one := zero.Inc()
	two := one.Inc()

	if two.Val() != 2 {
		t.Errorf("want: 2\ngot: %v", two.Val())
	}

	zero = two.Inc()
	if zero.Val() != 0 {
		t.Errorf("want: 0\ngot: %v", zero.Val())
	}

	two = two.SetMaxValue(3)
	three := two.Inc()
	if three.Val() != 3 {
		t.Errorf("want: 3\ngot: %v", three.Val())
	}
}

func TestSetMaxValue_IfNewMaxLessThanValueThenSetValueToZero(t *testing.T) {
	inc := New()

	inc = inc.Inc()
	inc = inc.Inc()
	inc = inc.Inc()

	inc = inc.SetMaxValue(2)

	if inc.Val() != 0 {
		t.Errorf("want: 0, got: %v", inc.Val())
	}
}
