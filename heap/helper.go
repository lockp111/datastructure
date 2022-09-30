package heap

type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

type Ordered interface {
	Signed | Unsigned | ~float32 | ~float64 | ~string
}

type Simple[T Ordered] struct {
	array  []T
	length int
	flat   bool
}

func NewMin[T Ordered](data ...T) Interface[T] {
	return &Simple[T]{array: data, length: len(data), flat: false}
}

func NewMax[T Ordered](data ...T) Interface[T] {
	return &Simple[T]{array: data, length: len(data), flat: true}
}

// Len is the number of elements in the collection.
func (s *Simple[T]) Len() int {
	return s.length
}

// Less reports whether the element with index i
// must sort before the element with index j.
func (s *Simple[T]) Less(i int, j int) bool {
	if s.flat {
		return s.array[i] > s.array[j]
	}
	return s.array[i] < s.array[j]
}

// Swap swaps the elements with indexes i and j.
func (s *Simple[T]) Swap(i int, j int) {
	s.array[i], s.array[j] = s.array[j], s.array[i]
}

// Push add x as element Len()
func (s *Simple[T]) Push(x T) {
	if s.length < len(s.array) {
		s.array[s.length] = x
	} else {
		s.array = append(s.array, x)
	}
	s.length++
}

// Pop remove and return element Len() - 1.
func (s *Simple[T]) Pop() (v T) {
	if s.length == 0 {
		return
	}
	s.length--
	return s.array[s.length]
}
