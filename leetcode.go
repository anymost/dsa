package main

//
//import (
//	"sort"
//	"strconv"
//	"strings"
//)
//
//type Person struct {
//	name string
//}
//
//type Singer interface {
//	sayName() string
//}
//
//type ListNode struct {
//	Val  int
//	Next *ListNode
//}
//
//func addDigits(num int) int {
//	if num < 10 {
//		return num
//	}
//	numList := strings.Split(strconv.Itoa(num), "")
//	count := 0
//	for _, val := range numList {
//		item, _ := strconv.Atoi(val)
//		count += item
//	}
//	return addDigits(count)
//}
//
//func sortArrayByParity(A []int) []int {
//	for i := 0; i < len(A); i++ {
//		if A[i]%2 == 1 {
//			B := append(A[0:i], A[i+1:]...)
//			A = append(B, A[i])
//		}
//	}
//	return A
//}
//
//func isOneBitCharacter(bits []int) bool {
//	for len(bits) > 2 {
//		if bits[0] == 0 {
//			bits = bits[1:]
//		} else {
//			bits = bits[2:]
//		}
//	}
//	return bits[0] == 0
//}
//
//func addBinary(a string, b string) string {
//	aList := strings.Split(a, "")
//	bList := strings.Split(b, "")
//	lenA := len(aList)
//	lenB := len(bList)
//	maxLen := 0
//	val := ""
//	if lenB > lenA {
//		maxLen = lenB
//		for i := 0; i < lenB-lenA; i++ {
//			aList = append([]string{"0"}, aList...)
//		}
//	} else {
//		maxLen = lenA
//		for i := 0; i < lenA-lenB; i++ {
//			bList = append([]string{"0"}, bList...)
//		}
//	}
//	result := make([]int, 0)
//	temp := 0
//	i := maxLen - 1
//	for i > -1 {
//		valA, _ := strconv.Atoi(aList[i])
//		valB, _ := strconv.Atoi(bList[i])
//		if valA^valB == 1 {
//			if temp == 0 {
//				result = append([]int{1}, result...)
//				val = "1" + val
//
//			} else {
//				val = "0" + val
//				temp = 1
//			}
//		} else if valA&valB == 1 {
//			if temp == 0 {
//				val = "0" + val
//				temp = 1
//			} else {
//				val = "1" + val
//				temp = 1
//			}
//		} else {
//			if temp == 0 {
//				val = "0" + val
//			} else {
//				val = "1" + val
//				temp = 0
//			}
//		}
//		i--
//	}
//	if temp == 1 {
//		val = "1" + val
//	}
//	return val
//}
//
//func addStrings(num1 string, num2 string) string {
//	aList, bList := strings.Split(num1, ""), strings.Split(num2, "")
//	lenA, lenB, maxLen := len(aList), len(bList), 0
//	if lenB > lenA {
//		maxLen = lenB
//		for i := 0; i < lenB-lenA; i++ {
//			aList = append([]string{"0"}, aList...)
//		}
//	} else {
//		maxLen = lenA
//		for i := 0; i < lenA-lenB; i++ {
//			bList = append([]string{"0"}, bList...)
//		}
//	}
//	val, temp, i := "", 0, maxLen-1
//	for i > -1 {
//		valA, _ := strconv.Atoi(aList[i])
//		valB, _ := strconv.Atoi(bList[i])
//		sum := valA + valB + temp
//		if sum >= 10 {
//			val = strconv.Itoa(sum%10) + val
//			temp = 1
//		} else {
//			val = strconv.Itoa(sum) + val
//			temp = 0
//		}
//		i--
//	}
//	if temp == 1 {
//		val = "1" + val
//	}
//	return val
//}
//
//func addToArrayForm(A []int, K int) []int {
//	B, i, temp := make([]int, len(A)), len(A)-1, 0
//	for K > 10 {
//		B[i] = K % 10
//		K = K / 10
//		i--
//	}
//	B[i] = K
//	for i := len(A) - 1; i > -1; i-- {
//		sum := A[i] + B[i] + temp
//		if sum >= 10 {
//			A[i] = sum % 10
//			temp = 1
//		} else {
//			A[i] = sum
//			temp = 0
//		}
//	}
//	if temp == 1 {
//		A = append([]int{1}, A...)
//	}
//	return A
//}
//
//func containsDuplicate(nums []int) bool {
//	sort.Ints(nums)
//	for i := 0; i < len(nums)-1; i++ {
//		if nums[i] == nums[i+1] {
//			return true
//		}
//	}
//	return false
//}
//
//func quickSort(array []int) {
//	if len(array) <= 1 {
//		return
//	}
//	mid, i := array[0], 1
//	head, tail := 0, len(array)-1
//	for head < tail {
//		if array[i] > mid {
//			array[i], array[tail] = array[tail], array[i]
//			tail--
//		} else {
//			array[i], array[head] = array[head], array[i]
//			head++
//			i++
//		}
//	}
//	array[head] = mid
//	quickSort(array[:head])
//	quickSort(array[head+1:])
//}
//
//func arrayPairSum(nums []int) int {
//	quickSort(nums)
//	sum := 0
//	for i := 0; i < len(nums); i += 2 {
//		sum += nums[i]
//	}
//	return sum
//}
//
//func hasAlternatingBits(n int) bool {
//	array := make([]int, 0)
//	for n != 0 {
//		temp := n % 2
//		array = append(array, temp)
//		n = (n - temp) / 2
//	}
//	for i := 0; i < len(array)-1; i++ {
//		if array[i] == array[i+1] {
//			return false
//		}
//	}
//	return true
//}
//
//func deleteNode(node *ListNode) {
//	node.Val = node.Next.Val
//	node.Next = node.Next.Next
//}
//
//func bitwiseComplement(N int) int {
//	num, val, temp := 0, 0, 0
//	for N != 0 {
//		temp = N % 2
//		N = N / 2
//		val = temp ^ 1
//		num = num<<1 + val
//	}
//	return num
//}
//
//type MyHashMap struct {
//	value []int
//}
//
///** Initialize your data structure here. */
//func Constructor() MyHashMap {
//	value := make([]int, 1000000)
//	for index := range value {
//		value[index] = -1
//	}
//	return MyHashMap{
//		value,
//	}
//}
//
///** value will always be non-negative. */
//func (this *MyHashMap) Put(key int, value int) {
//	this.value[key] = value
//}
//
///** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
//func (this *MyHashMap) Get(key int) int {
//	return this.value[key]
//}
//
///** Removes the mapping of the specified value key if this map contains a mapping for the key */
//func (this *MyHashMap) Remove(key int) {
//	this.value[key] = -1
//}
//
//func findShortestSubArray(nums []int) int {
//	count := make(map[int]int, 0)
//	for _, item := range nums {
//		if val, isExit := count[item]; isExit {
//			count[item] = val + 1
//		} else {
//			count[item] = 1
//		}
//	}
//	largest := 0
//	for _, val := range count {
//		if val > largest {
//			largest = val
//		}
//	}
//	result := 0
//	for _, val := range count {
//		if val == largest {
//			result++
//		}
//	}
//	return result
//}

//func search(nums []int, target int) int {
//	low, hi := 0, len(nums)
//	for low < hi {
//		mi := (low + hi) >> 1
//		if nums[mi] < target {
//			low = mi + 1
//		} else if target < nums[mi] {
//			hi = mi
//		} else {
//			return mi
//		}
//	}
//	return -1
//}

//func findLengthOfLCIS(nums []int) int {
//	length := len(nums)
//	if length == 0 {
//		return 0
//	}
//	counts := make([]int, 0, 10)
//	count := 1
//	for i := 0; i < length-1; i++ {
//		if nums[i] < nums[i+1] {
//			count++
//		} else {
//			counts = append(counts, count)
//			count = 1
//		}
//	}
//	counts = append(counts, count)
//	max := 1
//	for _, val := range counts {
//		if val > max {
//			max = val
//		}
//	}
//	return max
//}
//
//type MyLinkedList struct {
//	val []int
//}
//
//
///** Initialize your data structure here. */
//func Constructor() MyLinkedList {
//	return MyLinkedList{
//		val: make([]int, 0),
//	}
//}
//
//
///** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
//func (this *MyLinkedList) Get(index int) int {
//	if index < 0 || index >= len(this.val) || len(this.val) == 0 {
//		return -1
//	}
//	return this.val[index]
//}
//
//
///** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
//func (this *MyLinkedList) AddAtHead(val int)  {
//	this.val = append([]int{val}, this.val...)
//}
//
//
///** Append a node of value val to the last element of the linked list. */
//func (this *MyLinkedList) AddAtTail(val int)  {
//	this.val = append(this.val, val)
//}
//
//
///** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
//func (this *MyLinkedList) AddAtIndex(index int, val int)  {
//	if index == -1 {
//		index = 0
//	}
//	if index >= 0 && index <= len(this.val) {
//		if index == len(this.val) {
//			this.AddAtTail(val)
//		} else {
//			array := make([]int, 0)
//			array = append(array, this.val...)
//			array = append(array[0: index], val)
//			this.val = append(array, this.val[index:]...)
//		}
//	}
//}
//
//
///** Delete the index-th node in the linked list, if the index is valid. */
//func (this *MyLinkedList) DeleteAtIndex(index int)  {
//	if index >= 0 && len(this.val) > index {
//		this.val = append(this.val[0:index], this.val[index+1:]...)
//	}
//}

//func match(wordList []string, words string) bool {
//	for _, w := range wordList {
//		if !strings.Contains(words, w) {
//			return false
//		}
//	}
//	return true
//}
//
//func findWords(words []string) []string {
//	array := []string{"qwertyuiop", "asdfghjkl", "zxcvbnm"}
//	result := make([]string, 0)
//	for _, word := range words {
//		wordList := strings.Split(strings.ToLower(word), "")
//		if match(wordList, array[0]) || match(wordList, array[1]) || match(wordList, array[2]) {
//			result = append(result, word)
//		}
//	}
//	return result
//}