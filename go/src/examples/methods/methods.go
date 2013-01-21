package methods

type rect struct {
	width, height int
}

func (r rect) area() int {
	return r.width * r.height
}

func (r rect) perimeter() int {
	return 2 * (r.width + r.height)
}

// Always take by pointer to mutate values.
func (r *rect) mutate(w int, h int) {
	r.width = w
	r.height = h
}

// This takes r by value. So any mutation is local.
func (r rect) dummyMutate(w int, h int) {
	r.width = w
	r.height = h
}
