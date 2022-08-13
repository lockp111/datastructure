package internal

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
func bubbleSort(arr []string) {
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
	bubbleSort(arr)
	t.Log("-----------------after------------------")
	output(t, arr)
}

// 切分函数，并返回切分元素的下标
func partition(array []string, begin, end int) int {
	i := begin + 1 // 将array[begin]作为基准数，因此从array[begin+1]开始与基准数比较！
	j := end       // array[end]是数组的最后一位

	// 没重合之前
	for i < j {
		if len(array[i]) > len(array[begin]) {
			array[i], array[j] = array[j], array[i] // 交换
			j--
		} else {
			i++
		}
	}

	/* 跳出while循环后，i = j。
	 * 此时数组被分割成两个部分  -->  array[begin+1] ~ array[i-1] < array[begin]
	 *                        -->  array[i+1] ~ array[end] > array[begin]
	 * 这个时候将数组array分成两个部分，再将array[i]与array[begin]进行比较，决定array[i]的位置。
	 * 最后将array[i]与array[begin]交换，进行两个分割部分的排序！以此类推，直到最后i = j不满足条件就退出！
	 */
	if len(array[i]) >= len(array[begin]) { // 这里必须要取等“>=”，否则数组元素由相同的值组成时，会出现错误！
		i--
	}

	array[begin], array[i] = array[i], array[begin]
	return i
}

// 普通快速排序
func quickSort(array []string, begin, end int) {
	if begin < end {
		// 进行切分
		loc := partition(array, begin, end)
		// 对左部分进行快排
		quickSort(array, begin, loc-1)
		// 对右部分进行快排
		quickSort(array, loc+1, end)
	}
}

func TestQuick(t *testing.T) {
	const count = 10
	arr := newGraph(count)
	output(t, arr)
	quickSort(arr, 0, len(arr)-1)
	t.Log("-----------------after------------------")
	output(t, arr)
}
