// Package datastructures provides a collection of important data structures and algorithms,
// along with demonstration functions for several standard library packages.
package datastructures

import (
	"container/heap"
	"container/list"
	"container/ring"
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// ======================================================
// Linked List Implementation (Singly Linked List)
// ======================================================

// ListNode represents a node in a singly linked list.
type ListNode struct {
	Data interface{}
	Next *ListNode
}

// LinkedList represents a singly linked list.
type LinkedList struct {
	Head *ListNode
}

// Append adds a new node with the given data at the end of the list.
func (l *LinkedList) Append(data interface{}) {
	newNode := &ListNode{Data: data}
	if l.Head == nil {
		l.Head = newNode
		return
	}
	current := l.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = newNode
}

// Prepend adds a new node with the given data at the beginning of the list.
func (l *LinkedList) Prepend(data interface{}) {
	newNode := &ListNode{Data: data, Next: l.Head}
	l.Head = newNode
}

// Delete removes the first node containing the given data.
func (l *LinkedList) Delete(data interface{}) {
	if l.Head == nil {
		return
	}
	if l.Head.Data == data {
		l.Head = l.Head.Next
		return
	}
	current := l.Head
	for current.Next != nil {
		if current.Next.Data == data {
			current.Next = current.Next.Next
			return
		}
		current = current.Next
	}
}

// Print displays the linked list.
func (l *LinkedList) Print() {
	current := l.Head
	for current != nil {
		fmt.Printf("%v -> ", current.Data)
		current = current.Next
	}
	fmt.Println("nil")
}

// ======================================================
// Stack Implementation
// ======================================================

// Stack represents a stack data structure using a slice.
type Stack struct {
	items []interface{}
}

// Push adds an item onto the stack.
func (s *Stack) Push(item interface{}) {
	s.items = append(s.items, item)
}

// Pop removes and returns the top item from the stack. Returns false if empty.
func (s *Stack) Pop() (interface{}, bool) {
	if len(s.items) == 0 {
		return nil, false
	}
	index := len(s.items) - 1
	item := s.items[index]
	s.items = s.items[:index]
	return item, true
}

// Peek returns the top item without removing it. Returns false if empty.
func (s *Stack) Peek() (interface{}, bool) {
	if len(s.items) == 0 {
		return nil, false
	}
	return s.items[len(s.items)-1], true
}

// IsEmpty returns true if the stack is empty.
func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

// ======================================================
// Queue Implementation
// ======================================================

// Queue represents a queue data structure using a slice.
type Queue struct {
	items []interface{}
}

// Enqueue adds an item to the end of the queue.
func (q *Queue) Enqueue(item interface{}) {
	q.items = append(q.items, item)
}

// Dequeue removes and returns the item at the front of the queue. Returns false if empty.
func (q *Queue) Dequeue() (interface{}, bool) {
	if len(q.items) == 0 {
		return nil, false
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item, true
}

// Peek returns the item at the front of the queue without removing it. Returns false if empty.
func (q *Queue) Peek() (interface{}, bool) {
	if len(q.items) == 0 {
		return nil, false
	}
	return q.items[0], true
}

// IsEmpty returns true if the queue is empty.
func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}

// ======================================================
// Binary Search Tree Implementation
// ======================================================

// TreeNode represents a node in a binary search tree.
type TreeNode struct {
	Data  int
	Left  *TreeNode
	Right *TreeNode
}

// BinarySearchTree represents a binary search tree.
type BinarySearchTree struct {
	Root *TreeNode
}

// Insert inserts a new integer into the BST.
func (bst *BinarySearchTree) Insert(data int) {
	bst.Root = insertNode(bst.Root, data)
}

func insertNode(node *TreeNode, data int) *TreeNode {
	if node == nil {
		return &TreeNode{Data: data}
	}
	if data < node.Data {
		node.Left = insertNode(node.Left, data)
	} else {
		node.Right = insertNode(node.Right, data)
	}
	return node
}

// Search returns true if data exists in the BST.
func (bst *BinarySearchTree) Search(data int) bool {
	return searchNode(bst.Root, data)
}

func searchNode(node *TreeNode, data int) bool {
	if node == nil {
		return false
	}
	if node.Data == data {
		return true
	} else if data < node.Data {
		return searchNode(node.Left, data)
	} else {
		return searchNode(node.Right, data)
	}
}

// InOrderTraversal returns the in-order traversal of the BST as a slice.
func (bst *BinarySearchTree) InOrderTraversal() []int {
	var result []int
	inOrder(bst.Root, &result)
	return result
}

func inOrder(node *TreeNode, result *[]int) {
	if node != nil {
		inOrder(node.Left, result)
		*result = append(*result, node.Data)
		inOrder(node.Right, result)
	}
}

// ======================================================
// Priority Queue Implementation (Using container/heap)
// ======================================================

// PriorityQueueItem represents an item in the priority queue.
type PriorityQueueItem struct {
	Value    interface{}
	Priority int // Lower values indicate higher priority.
	Index    int // Needed by the heap.Interface methods.
}

// PriorityQueue implements heap.Interface and holds PriorityQueueItems.
type PriorityQueue []*PriorityQueueItem

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Priority < pq[j].Priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

// Push adds an item to the priority queue.
func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*PriorityQueueItem)
	item.Index = n
	*pq = append(*pq, item)
}

// Pop removes and returns the item with the highest priority.
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // Avoid memory leak
	item.Index = -1 // For safety
	*pq = old[0 : n-1]
	return item
}

// Update modifies the priority and value of an item in the queue.
func (pq *PriorityQueue) Update(item *PriorityQueueItem, value interface{}, priority int) {
	item.Value = value
	item.Priority = priority
	heap.Fix(pq, item.Index)
}

// ======================================================
// Binary Search Algorithm
// ======================================================

// BinarySearch performs a binary search on a sorted slice of integers.
// Returns the index of the target or -1 if not found.
func BinarySearch(arr []int, target int) int {
	low, high := 0, len(arr)-1
	for low <= high {
		mid := low + (high-low)/2
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

// ======================================================
// Additional Demonstration Functions for Standard Packages
// ======================================================

// DemoContainerList demonstrates the use of container/list (a doubly linked list).
func DemoContainerList() {
	// Create a new doubly linked list.
	l := list.New()
	l.PushBack(10)
	l.PushBack(20)
	l.PushBack(30)
	fmt.Print("container/list (doubly linked list): ")
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Printf("%v ", e.Value)
	}
	fmt.Println()
}

// DemoContainerHeap demonstrates the use of container/heap (priority queue) using our PriorityQueue.
func DemoContainerHeap() {
	// Create some tasks with priorities (lower value means higher priority).
	items := []*PriorityQueueItem{
		{Value: "Task1", Priority: 3},
		{Value: "Task2", Priority: 1},
		{Value: "Task3", Priority: 2},
	}
	// Initialize the priority queue.
	pq := make(PriorityQueue, len(items))
	for i, item := range items {
		pq[i] = item
		pq[i].Index = i
	}
	heap.Init(&pq)
	// Push an additional task with highest priority.
	heap.Push(&pq, &PriorityQueueItem{Value: "Task4", Priority: 0})
	// Pop the highest priority task.
	itemPopped := heap.Pop(&pq).(*PriorityQueueItem)
	fmt.Println("container/heap (priority queue) - Highest Priority Task:", itemPopped.Value)
}

// DemoContainerRing demonstrates the use of container/ring (circular list).
func DemoContainerRing() {
	// Create a ring with 5 elements.
	r := ring.New(5)
	// Fill the ring with sequential numbers.
	for i := 0; i < r.Len(); i++ {
		r.Value = i + 1
		r = r.Next()
	}
	// Iterate over the ring and print the values.
	fmt.Print("container/ring (circular list): ")
	r.Do(func(value interface{}) {
		fmt.Printf("%v ", value)
	})
	fmt.Println()
}

// DemoSort demonstrates sorting a slice using the sort package.
func DemoSort() {
	// Create an unsorted slice.
	a := []int{3, 1, 2}
	// Sort the slice in ascending order.
	sort.Ints(a)
	fmt.Println("sort.Ints (sorted slice):", a)
}

// DemoMathRand demonstrates generating a random number using math/rand.
func DemoMathRand() {
	// Seed the random number generator using the current time.
	rand.Seed(time.Now().UnixNano())
	// Generate a random number between 0 and 99.
	r := rand.Intn(100)
	fmt.Println("math/rand.Intn(100):", r)
}
