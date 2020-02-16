package counter

const MaxUint = ^uint(0) // set max value by reversing all bits

type Counter struct {
	counter uint
	maxVal  uint
}

// New creates a new counter started with zero and max value equals to MaxUint
func New() Counter {
	return Counter {
		maxVal: MaxUint,
	}
}

// MaxVal returns the max value
func (c Counter) MaxVal() uint {
	return c.maxVal
}

// Val returns the current value
func (c Counter) Val() uint {
	return c.counter
}

// Inc increments the counter and returns a new Counter.
// Inc increments the value until it becomes the max value, after the next increment,
// the value will be set to zero
func (c Counter) Inc() Counter {
	if c.counter < c.maxVal {
		c.counter++
	} else {
		c.counter = 0
	}
	return c
}

// SetMaxValue sets a new max value and returns a new Counter.
// If the new max value is less than current value, then counter drops its value
func (c Counter) SetMaxValue(n uint) Counter {
	if n < c.counter {
		c.counter = 0
	}

	c.maxVal = n
	return c
}
