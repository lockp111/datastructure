package hashring

import (
	"hash/crc32"
	"sort"
)

type Hashring[T ISlot] struct {
	slotMap map[uint32]Slot[T]
	slots   []uint32
}

// sort
func (h *Hashring[T]) Len() int           { return len(h.slots) }
func (h *Hashring[T]) Less(i, j int) bool { return h.slots[i] < h.slots[j] }
func (h *Hashring[T]) Swap(i, j int)      { h.slots[i], h.slots[j] = h.slots[j], h.slots[i] }

func Hash(key string) uint32 {
	return crc32.ChecksumIEEE([]byte(key))
}

func New[T ISlot]() *Hashring[T] {
	return &Hashring[T]{slotMap: make(map[uint32]Slot[T]), slots: make([]uint32, 0, 2048)}
}

func (h *Hashring[T]) find(value uint32) int {
	count := h.Count()
	index := sort.Search(count, func(i int) bool {
		return h.slots[i] >= value
	})
	if index == count {
		return 0
	}
	return index
}

func (h *Hashring[T]) get(index int) (Slot[T], bool) {
	slot, ok := h.slotMap[h.slots[index]]
	if !ok {
		return slot, ok
	}
	slot.index = index
	return slot, true
}

// UnsortAdd slots to hashring, but no sort
func (h *Hashring[T]) UnsortAdd(slots ...Slot[T]) {
	for _, slot := range slots {
		_, ok := h.slotMap[slot.hash]
		if ok {
			continue
		}

		h.slotMap[slot.hash] = slot
		h.slots = append(h.slots, slot.hash)
	}
}

// Add slots to hashring. It will be slow if you call too often
func (h *Hashring[T]) Add(slots ...Slot[T]) {
	h.UnsortAdd(slots...)
	h.Sort()
}

// Sort hashring
func (h *Hashring[T]) Sort() {
	sort.Sort(h)
}

// Count slot number
func (h *Hashring[T]) Count() int {
	return len(h.slots)
}

// Get slot by key
func (h *Hashring[T]) Get(key string) (Slot[T], bool) {
	if len(h.slots) == 0 {
		return Slot[T]{}, false
	}

	hashed := Hash(key)
	index := h.find(hashed)
	return h.get(index)
}

// Remove slot by key
func (h *Hashring[T]) Remove(key string) {
	hashed := Hash(key)
	_, ok := h.slotMap[hashed]
	if !ok {
		return
	}
	if len(h.slots) == 1 {
		h.slotMap = make(map[uint32]Slot[T])
		h.slots = make([]uint32, 0, 2048)
		return
	}

	index := h.find(hashed)
	h.slots = append(h.slots[:index], h.slots[index+1:]...)
	delete(h.slotMap, hashed)
}

// BatchRemove slots
func (h *Hashring[T]) BatchRemove(keys []string) {
	begin := 0
	for _, key := range keys {
		hashed := Hash(key)
		_, ok := h.slotMap[hashed]
		if !ok {
			continue
		}
		if len(h.slots) == 1 {
			h.slotMap = make(map[uint32]Slot[T])
			h.slots = make([]uint32, 0, 2048)
			return
		}

		delete(h.slotMap, hashed)
		index := h.find(hashed)
		h.slots[begin], h.slots[index] = h.slots[index], h.slots[begin]
		begin++
	}
	h.slots = h.slots[begin:]
	sort.Sort(h)
}

// ForEach hashring
func (h *Hashring[T]) ForEach(fn func(index int, hash uint32, value T)) {
	for index, hashed := range h.slots {
		slot := h.slotMap[hashed]
		fn(index, hashed, slot.value)
	}
}

// GetNext from slot
func (h *Hashring[T]) GetNext(s Slot[T]) Slot[T] {
	index := s.index + 1
	if index >= len(h.slots) {
		index = 0
	}

	slot, _ := h.get(index)
	return slot
}

// GetPrev from slot
func (h *Hashring[T]) GetPrev(s Slot[T]) Slot[T] {
	index := s.index - 1
	if index < 0 {
		index = len(h.slots) - 1
	}

	slot, _ := h.get(index)
	return slot
}
