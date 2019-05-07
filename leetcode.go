package main

import (
	"fmt"
	"regexp"
)

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

//
//func removeOuterParentheses(S string) string {
//	list := strings.Split(S, "")
//	stack := NewStack()
//	removeIndex := make([]int, 0)
//	type Item struct {
//		index int
//		val string
//	}
//	for i, v := range list {
//		if v == "(" {
//			stack.Push(&Item{
//				index: i,
//				val: v,
//			})
//		} else {
//			top := stack.Pop()
//			if !stack.Empty() {
//				removeIndex  = append(removeIndex, i)
//				if val, ok := top.(*Item); ok {
//					removeIndex = append(removeIndex, val.index)
//				}
//			}
//		}
//	}
//	sort.Ints(removeIndex)
//	for i := len(removeIndex) - 1; i > -1; i++ {
//		v := removeIndex[i]
//		list = append(list[0: v], list[v+1:]...)
//	}
//	return strings.Join(list, "")
//}

//func reverseWords(s string) string {
//	list := strings.Split(s, " ")
//
//	for index, val := range list {
//		wordList := strings.Split(val, "")
//		length := len(wordList)
//		for index := 0; index < length/2; index++ {
//			if index < length/2 {
//				wordList[index], wordList[length-1-index] = wordList[length-1-index], wordList[index]
//			}
//		}
//		list[index] = strings.Join(wordList, "")
//	}
//	return strings.Join(list, " ")
//}

//func reverseString(s []byte) {
//	length := len(s)
//	for i := 0; i < length/2; i++ {
//		s[i], s[length-1-i] = s[length-1-i], s[i]
//	}
//}
//
//func reverse(s []string) {
//	length := len(s)
//	for i := range s {
//		if i < length/2 {
//			s[i], s[length-1-i] = s[length-1-i], s[i]
//		}
//	}
//}
//
//func reverseStr(s string, k int) string {
//	length := len(s)
//	if k == 1 {
//		return s
//	}
//	list := strings.Split(s, "")
//	if length <= k {
//		reverse(list)
//	} else {
//		for i := 0; i < length; {
//			if i+k > length {
//				reverse(list[i : length-1])
//			} else {
//				reverse(list[i : i+k])
//			}
//			i += 2 * k
//
//		}
//	}
//	return strings.Join(list, "")
//}

//func thirdMax(nums []int) int {
//	first, second, third := math.MinInt64, math.MinInt64, math.MinInt64
//	for _, val := range nums {
//		if val > third {
//			if val > second {
//				if val > first {
//					third = second
//					second = first
//					first = val
//				} else if val < first{
//					third = second
//					second = val
//				}
//			} else if val < second {
//				third = val
//			}
//		}
//	}
//
//	if  third == math.MinInt64 {
//		return first
//	} else {
//		return third
//	}
//}
//
//func twoSum(numbers []int, target int) []int {
//	for i, v1 := range numbers {
//		for j, v2 := range numbers {
//			if i != j && v1 + v2 == target {
//				return []int{i, j}
//			}
//		}
//	}
//	return []int{}
//}

//func transpose(A [][]int) [][]int {
//	externalLen, internalLen := len(A), len(A[0])
//	matrix := make([][]int, internalLen)
//	for i := range matrix {
//		matrix[i] = make([]int, externalLen)
//	}
//	for i := 0; i < externalLen; i++ {
//		for k := 0; k < internalLen; k++ {
//			matrix[k][i] = A[i][k]
//		}
//	}
//	return matrix
//}
//
//func isToeplitzMatrix(matrix [][]int) bool {
//	rowLen := len(matrix[0])
//	sum := func(array []int) int {
//		sum := 0
//		for _, val := range array {
//			sum += val
//		}
//		return sum
//	}
//	for i := 0; i < len(matrix)-1; i++ {
//		prev := matrix[i]
//		next := matrix[i+1]
//		if sum(prev[0:rowLen-1]) != sum(next[1:]) {
//			return false
//		}
//	}
//	return true
//}

//func validMountainArray(A []int) bool {
//	arrayLen := len(A)
//	if arrayLen < 3 {
//		return false
//	}
//	maxIndex, max := -1, -1
//	for index, val := range A {
//		if val > max {
//			max, maxIndex = val, index
//		}
//	}
//	if maxIndex == arrayLen-1 {
//		return false
//	}
//	for i := 0; i < maxIndex-1; i++ {
//		if A[i] >= A[i+1] {
//			return false
//		}
//	}
//	for i := maxIndex; i < arrayLen-1; i++ {
//		if A[i] <= A[i+1] {
//			return false
//		}
//	}
//	return true
//}

func hasGroupsSizeX(deck []int) bool {
	if len(deck) < 2 {
		return false
	}
	countMap := make(map[int]int)
	for _, v := range deck {
		if val, isExit := countMap[v]; isExit {
			countMap[v] = val + 1
		} else {
			countMap[v] = 1
		}
	}
	countList := make([]int, 0)
	if len(countMap) < 2  {
		return false
	}
	for _, v := range countMap {
		if v == 1  {
			return false
		}
		countList = append(countList, v)
	}
	if len(countList) == 1 {
		return true
	}
	min := 1
	commonChild := func(a int, b int) int {
		for a != b {
			if b < a {
				a -= b
			} else if a < b {
				b -= a
			}
		}
		return a
	}
	for i := 0; i < len(countList)-1; i++ {
		min = commonChild(countList[i], countList[i+1])
	}
	if min == 1 {
		return false
	} else {
		return true
	}
}

func main() {
	//val := hasGroupsSizeX([]int{1, 1, 1, 2, 2, 2, 3, 3})
	//fmt.Println(val)
	isMatch, _ := regexp.MatchString("^[a-z]+$", "abcdefg12343")
	pattern, _ := regexp.Compile("^[a-z]+")
	val := pattern.FindString("abcd123")
	fmt.Println(val)
	fmt.Println(isMatch)
}
