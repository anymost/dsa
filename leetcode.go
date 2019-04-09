package main

import (
	"sort"
	"strconv"
	"strings"
)

type Person struct {
	name string
}

type Singer interface {
	sayName() string
}


type ListNode struct {
	Val  int
	Next *ListNode
}

func addDigits(num int) int {
	if num < 10 {
		return num
	}
	numList := strings.Split(strconv.Itoa(num), "")
	count := 0
	for _, val := range numList {
		item, _ := strconv.Atoi(val)
		count += item
	}
	return addDigits(count)
}

func sortArrayByParity(A []int) []int {
	for i := 0; i < len(A); i++ {
		if A[i]%2 == 1 {
			B := append(A[0:i], A[i+1:]...)
			A = append(B, A[i])
		}
	}
	return A
}

func isOneBitCharacter(bits []int) bool {
	for len(bits) > 2 {
		if bits[0] == 0 {
			bits = bits[1:]
		} else {
			bits = bits[2:]
		}
	}
	return bits[0] == 0
}

func addBinary(a string, b string) string {
	aList := strings.Split(a, "")
	bList := strings.Split(b, "")
	lenA := len(aList)
	lenB := len(bList)
	maxLen := 0
	val := ""
	if lenB > lenA {
		maxLen = lenB
		for i := 0; i < lenB-lenA; i++ {
			aList = append([]string{"0"}, aList...)
		}
	} else {
		maxLen = lenA
		for i := 0; i < lenA-lenB; i++ {
			bList = append([]string{"0"}, bList...)
		}
	}
	result := make([]int, 0)
	temp := 0
	i := maxLen - 1
	for i > -1 {
		valA, _ := strconv.Atoi(aList[i])
		valB, _ := strconv.Atoi(bList[i])
		if valA^valB == 1 {
			if temp == 0 {
				result = append([]int{1}, result...)
				val = "1" + val

			} else {
				val = "0" + val
				temp = 1
			}
		} else if valA&valB == 1 {
			if temp == 0 {
				val = "0" + val
				temp = 1
			} else {
				val = "1" + val
				temp = 1
			}
		} else {
			if temp == 0 {
				val = "0" + val
			} else {
				val = "1" + val
				temp = 0
			}
		}
		i--
	}
	if temp == 1 {
		val = "1" + val
	}
	return val
}

func addStrings(num1 string, num2 string) string {
	aList, bList := strings.Split(num1, ""), strings.Split(num2, "")
	lenA, lenB, maxLen := len(aList), len(bList), 0
	if lenB > lenA {
		maxLen = lenB
		for i := 0; i < lenB-lenA; i++ {
			aList = append([]string{"0"}, aList...)
		}
	} else {
		maxLen = lenA
		for i := 0; i < lenA-lenB; i++ {
			bList = append([]string{"0"}, bList...)
		}
	}
	val, temp, i := "", 0, maxLen-1
	for i > -1 {
		valA, _ := strconv.Atoi(aList[i])
		valB, _ := strconv.Atoi(bList[i])
		sum := valA + valB + temp
		if sum >= 10 {
			val = strconv.Itoa(sum%10) + val
			temp = 1
		} else {
			val = strconv.Itoa(sum) + val
			temp = 0
		}
		i--
	}
	if temp == 1 {
		val = "1" + val
	}
	return val
}

func addToArrayForm(A []int, K int) []int {
	B, i, temp := make([]int, len(A)), len(A)-1, 0
	for K > 10 {
		B[i] = K % 10
		K = K / 10
		i--
	}
	B[i] = K
	for i := len(A) - 1; i > -1; i-- {
		sum := A[i] + B[i] + temp
		if sum >= 10 {
			A[i] = sum % 10
			temp = 1
		} else {
			A[i] = sum
			temp = 0
		}
	}
	if temp == 1 {
		A = append([]int{1}, A...)
	}
	return A
}

func containsDuplicate(nums []int) bool {
	sort.Ints(nums)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			return true
		}
	}
	return false
}

func quickSort(array []int) {
	if len(array) <= 1 {
		return
	}
	mid, i := array[0], 1
	head, tail := 0, len(array)-1
	for head < tail {
		if array[i] > mid {
			array[i], array[tail] = array[tail], array[i]
			tail--
		} else {
			array[i], array[head] = array[head], array[i]
			head++
			i++
		}
	}
	array[head] = mid
	quickSort(array[:head])
	quickSort(array[head+1:])
}

func arrayPairSum(nums []int) int {
	quickSort(nums)
	sum := 0
	for i := 0; i < len(nums); i += 2 {
		sum += nums[i]
	}
	return sum
}

func hasAlternatingBits(n int) bool {
	array := make([]int, 0)
	for n != 0 {
		temp := n % 2
		array = append(array, temp)
		n = (n - temp) / 2
	}
	for i := 0; i < len(array)-1; i++ {
		if array[i] == array[i+1] {
			return false
		}
	}
	return true
}

func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}

func bitwiseComplement(N int) int {
	num, val, temp := 0, 0, 0
	for N != 0 {
		temp = N % 2
		N = N / 2
		val = temp ^ 1
		num = num << 1 + val
	}
	return num
}


