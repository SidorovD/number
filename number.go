package number

const (
	Max   = 4294967295 // use 32bit uint max value for portability
)

type Number struct {
	n   uint
	max uint
}

func New() Number {
	return Number{
		n:   0,
		max: Max,
	}
}

func (i Number) Val() uint {
	return i.n
}

func (i Number) Inc() Number {
	if i.n < i.max {
		i.n++
	} else {
		i.n = 0
	}

	return i
}

func (i Number) SetMaxValue(n uint) Number {
	if n < i.n {
		i.n = 0
	}

	i.max = n
	return i
}
