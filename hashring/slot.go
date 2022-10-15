package hashring

type ISlot interface {
	Key() string
}

type Slot[T ISlot] struct {
	value T
	hash  uint32
	index int
}

func NewSlot[T ISlot](value T) Slot[T] {
	return Slot[T]{
		value: value,
		hash:  Hash(value.Key()),
	}
}

func (s *Slot[T]) GetValue() T {
	return s.value
}

func (s *Slot[T]) Hash() uint32 {
	return s.hash
}
