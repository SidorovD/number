package number

const Max = ^uint(0) // set max value by reversing all bits

type Number struct {
	n   uint
	max uint
}

// New creates a new number with a zero val and max value equal to the max value of uint
func New() Number {
	return Number{
		n:   0,
		max: Max,
	}
}

// MaxVal returns the max value
func (i Number) MaxVal() uint {
	return i.max
}

// Val returns the value
func (i Number) Val() uint {
	return i.n
}

// Inc increments a value and returns a new Number.
// Inc increments the value until it becomes the max value, after the next increment, 
// the value will be set to zero
func (i Number) Inc() Number {
	if i.n < i.max {
		i.n++
	} else {
		i.n = 0
	}
	return i
}

// SetMaxValue sets a new max value and returns a new Number.
// If the new max value is less than current value, then n becomes zero
func (i Number) SetMaxValue(n uint) Number {
	if n < i.n {
		i.n = 0
	}

	i.max = n
	return i
}

