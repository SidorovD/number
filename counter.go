package counter

const MaxUint = ^uint(0) // set max uint value by reversing all bits

// Counter is imutable one-directional overflowable counter. After it reaches its upper bound (maxVal), counter will reset its value
type Counter struct {
	counter uint
	maxVal  uint
}

// New creates a new counter that started from zero and max value that equals to MaxUint
func New() Counter {
	return Counter{
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
// The counter is incremented until it reaches the max value.
// After this, the next Inc will "overflow" the counter i.e. reset its value
func (c Counter) Inc() Counter {
	if c.counter < c.maxVal {
		c.counter++
	} else {
		c.counter = 0
	}
	return c
}

// SetMaxValue sets a new max value and returns a new Counter.
// If the new max value is less than current value, then counter resets
func (c Counter) SetMaxValue(n uint) Counter {
	if n < c.counter {
		c.counter = 0
	}
	c.maxVal = n
	return c
}
