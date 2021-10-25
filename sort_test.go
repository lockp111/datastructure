package datastructure

import (
	"crypto/rand"
	"math/big"
	"testing"
)

func newGraph(n int) []string {
	graphArr := make([]string, n)
	for i := range graphArr {
		n, _ := rand.Int(rand.Reader, big.NewInt(100))

		var (
			num = int(n.Int64())
			out string
		)
		for j := 0; j < num; j++ {
			out += "="
		}
		graphArr[i] = out
	}
	return graphArr
}

func output(t *testing.T, arr []string) {
	for _, v := range arr {
		t.Log(v)
	}
}

// 冒泡排序法
func sortBubble(arr []string) {
	count := len(arr)
	if count == 0 {
		return
	}
	for i := 0; i < count; i++ {
		for j := i + 1; j < count; j++ {
			if len(arr[i]) > len(arr[j]) {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}

func TestBubble(t *testing.T) {
	const count = 10
	arr := newGraph(count)
	output(t, arr)
	sortBubble(arr)
	t.Log("-----------------after------------------")
	output(t, arr)
}
