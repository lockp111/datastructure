package hashring_test

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"sort"
	"testing"

	"github.com/lockp111/datastructure/hashring"
)

type node struct {
	key   string
	value int
}

func (n node) Key() string {
	return n.key
}

var nodeList = []hashring.Slot[node]{
	hashring.NewSlot(node{key: "1", value: 1}),
	hashring.NewSlot(node{key: "2", value: 2}),
	hashring.NewSlot(node{key: "3", value: 3}),
	hashring.NewSlot(node{key: "4", value: 4}),
	hashring.NewSlot(node{key: "5", value: 5}),
	hashring.NewSlot(node{key: "6", value: 6}),
}

var sortNodes = []node{
	{key: "2", value: 2},
	{key: "6", value: 6},
	{key: "3", value: 3},
	{key: "1", value: 1},
	{key: "5", value: 5},
	{key: "4", value: 4},
}

func TestHashringGet(t *testing.T) {
	ring := hashring.New[node]()
	ring.Add(nodeList...)

	slot, ok := ring.Get("2")
	if !ok {
		t.Fatal()
	}

	if slot.GetValue().Key() != nodeList[1].GetValue().Key() {
		t.Fatal()
	}

	slot, _ = ring.Get("7")
	if slot.Hash() != nodeList[2].Hash() {
		t.Fatal()
	}
}

func TestHashringRemove(t *testing.T) {
	ring := hashring.New[node]()
	ring.Add(nodeList...)

	slot, ok := ring.Get("4")
	if !ok {
		t.Fatal()
	}
	if slot.Hash() != nodeList[3].Hash() {
		t.Fatal()
	}

	ring.Remove("4")

	slot, _ = ring.Get("4")
	if slot.Hash() != nodeList[1].Hash() {
		t.Fatal()
	}
}

func TestUnsortAdd(t *testing.T) {
	ring := hashring.New[node]()
	ring.UnsortAdd(nodeList...)
	ring.Sort()

	ring.ForEach(func(index int, hash uint32, value node) {
		node := sortNodes[index]
		if node != value {
			t.Fatal()
		}
	})
}

func TestBatchRemove(t *testing.T) {
	ring := hashring.New[node]()
	for _, n := range nodeList {
		ring.Add(n)
	}

	ring.BatchRemove([]string{"1", "2", "3"})

	var wantNodes = []node{
		{key: "6", value: 6},
		{key: "5", value: 5},
		{key: "4", value: 4},
	}
	ring.ForEach(func(index int, hash uint32, value node) {
		node := wantNodes[index]
		if node != value {
			t.Fatal()
		}
	})
}

func TestGetNext(t *testing.T) {
	ring := hashring.New[node]()
	for _, n := range nodeList {
		ring.Add(n)
	}

	slot, ok := ring.Get("3")
	if !ok {
		t.Fatal()
	}
	if slot.GetValue() != sortNodes[2] {
		t.Fatal()
	}

	slot = ring.GetNext(slot)
	if slot.GetValue() != sortNodes[3] {
		t.Fatal()
	}

	index := 3
	for i := 0; i < ring.Count(); i++ {
		index++
		if index >= ring.Count() {
			index = 0
		}
		slot = ring.GetNext(slot)
		if slot.GetValue() != sortNodes[index] {
			t.Fatal()
		}
	}
}

func TestGetPrev(t *testing.T) {
	ring := hashring.New[node]()
	for _, n := range nodeList {
		ring.Add(n)
	}

	slot, ok := ring.Get("3")
	if !ok {
		t.Fatal()
	}
	if slot.GetValue() != sortNodes[2] {
		t.Fatal()
	}

	slot = ring.GetPrev(slot)
	if slot.GetValue() != sortNodes[1] {
		t.Fatal()
	}

	index := 1
	for i := 0; i < ring.Count(); i++ {
		index--
		if index < 0 {
			index = ring.Count() - 1
		}
		slot = ring.GetPrev(slot)
		if slot.GetValue() != sortNodes[index] {
			t.Fatal()
		}
	}
}

func BenchmarkAdd(b *testing.B) {
	var slots []hashring.Slot[node]
	for j := 0; j < 1000000; j++ {
		n := node{fmt.Sprint(j), j}
		slots = append(slots, hashring.NewSlot(n))
	}

	b.Run("add one", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			ring := hashring.New[node]()
			for _, s := range slots {
				ring.UnsortAdd(s)
			}
			sort.Sort(ring)
		}
	})

	b.Run("batch add", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			ring := hashring.New[node]()
			ring.Add(slots...)
		}
	})

}

func BenchmarkRemove(b *testing.B) {
	ring := hashring.New[node]()
	var slots []hashring.Slot[node]
	for i := 0; i < 1000000; i++ {
		n := node{fmt.Sprint(i), i}
		slots = append(slots, hashring.NewSlot(n))
	}
	ring.Add(slots...)

	n := 500
	b.Run("remove", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			x := i * n
			for j := 0; j < n; j++ {
				ring.Remove(fmt.Sprint(j + x))
			}
		}
	})

	b.Run("batch remove", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			x := i * n
			s := make([]string, 0, 100)
			for j := 0; j < n; j++ {
				s = append(s, fmt.Sprint(j+x))
			}
			ring.BatchRemove(s)
		}
	})
}

func BenchmarkGet(b *testing.B) {
	ring := hashring.New[node]()
	var slots []hashring.Slot[node]
	for i := 0; i < 1000000; i++ {
		n := node{fmt.Sprint(i), i}
		slots = append(slots, hashring.NewSlot(n))
	}
	ring.Add(slots...)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		key, _ := rand.Int(rand.Reader, big.NewInt(1000000))
		ring.Get(key.String())
	}
}