package datastructure

/**
给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target 的那两个整数，并返回它们的数组下标。
你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。
你可以按任意顺序返回答案。

示例 1：
输入：nums = [2,7,11,15], target = 9
输出：[0,1]
解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1] 。

示例 2：
输入：nums = [3,2,4], target = 6
输出：[1,2]

示例 3：
输入：nums = [3,3], target = 6
输出：[0,1]

提示：
2 <= nums.length <= 104
-109 <= nums[i] <= 109
-109 <= target <= 109
只会存在一个有效答案

进阶：你可以想出一个时间复杂度小于 O(n2) 的算法吗？
*/
func twoSum1(nums []int, target int) []int {
	var out []int
	for i, v := range nums {
		value := target - v
		for j := i + 1; j < len(nums); j++ {
			if nums[j] != value {
				continue
			}
			out = append(out, i, j)
			return out
		}
	}
	return nil
}

func twoSum2(nums []int, target int) []int {
	var valueMap = make(map[int]int)
	for i, v := range nums {
		if vv, ok := valueMap[v]; ok {
			valueMap[v] = i
			i = vv
		} else {
			valueMap[v] = i
		}

		key := target - v
		if j, ok := valueMap[key]; ok && j != i {
			return append([]int{}, i, j)
		}
	}
	return nil
}

/*
给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。
请你将两个数相加，并以相同形式返回一个表示和的链表。

你可以假设除了数字 0 之外，这两个数都不会以 0 开头。

示例 1：
输入：l1 = [2,4,3], l2 = [5,6,4]
输出：[7,0,8]
解释：342 + 465 = 807.

示例 2：
输入：l1 = [0], l2 = [0]
输出：[0]

示例 3：
输入：l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
输出：[8,9,9,9,0,0,0,1]

提示：
每个链表中的节点数在范围 [1, 100] 内
0 <= Node.val <= 9
题目数据保证列表表示的数字不含前导零
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Next *ListNode
	Val  int
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var l1Arr, l2Arr []int
	for l1 != nil || l2 != nil {
		if l1 != nil {
			l1Arr = append(l1Arr, l1.Val)
			l1 = l1.Next
		}
		if l2 != nil {
			l2Arr = append(l2Arr, l2.Val)
			l2 = l2.Next
		}
	}

	var (
		i    int
		out  *ListNode
		last *ListNode
		more int
	)
	for len(l1Arr) > i || len(l2Arr) > i {
		var v1, v2 int
		if len(l1Arr) > i {
			v1 = l1Arr[i]
		}
		if len(l2Arr) > i {
			v2 = l2Arr[i]
		}
		i++

		value := (v1 + v2 + more) % 10
		more = (v1 + v2 + more) / 10

		if last == nil {
			last = &ListNode{
				Val: value,
			}
			out = last
			continue
		}
		last.Next = &ListNode{
			Val: value,
		}
		last = last.Next
	}

	if more > 0 {
		last.Next = &ListNode{
			Val: more,
		}
	}
	return out
}

/*
给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。

示例 1:
输入: s = "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。

示例 2:
输入: s = "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。

示例 3:
输入: s = "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。

示例 4:
输入: s = ""
输出: 0

提示：
0 <= s.length <= 5 * 104
s 由英文字母、数字、符号和空格组成
*/
func lengthOfLongestSubstring(s string) int {
	var (
		keys         = make(map[rune]int)
		left, length int
	)
	for right, v := range s {
		if i, ok := keys[v]; ok && i+1 > left {
			left = i + 1
		}

		keys[v] = right
		if right-left+1 > length {
			length = right - left + 1
		}
	}
	return length
}
