package main

import (
	"bufio"
	"container/heap"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/abtin81badie/GoLangEssentials/alias"
	"github.com/abtin81badie/GoLangEssentials/datastructures"
	"github.com/abtin81badie/GoLangEssentials/greeting"
	"github.com/abtin81badie/GoLangEssentials/mathutils"
	"github.com/abtin81badie/GoLangEssentials/stringutils"
)

/*
	Global Variables and Constants
	--------------------------------
	Note: Global variables and constants are generally discouraged for large projects,
	but are shown here for demonstration.
*/

// Basic iota
const (
	A = iota //0
	B        //1
	C        //2
	D        //3
)

// Using iota for custom step values
const (
	KB = 1 << (10 * iota) // 1 << (10 * 0) = 1
	MB                    // 1 << (10 * 1) = 1024
	GB                    // 1 << (10 * 2) = 1048576
	TB                    // 1 << (10 * 3) = 1073741824
)

type LogLevel int

const (
	DEBUG LogLevel = iota // 0
	INFO                  // 1
	WARN                  // 2
	ERROR                 // 3
)

// Global Variables Declaration
var (
	Age    int    // Global variable to store age (unused in this demo)
	Name   string // Global variable to store name (unused in this demo)
	Family string // Global variable to store family name (unused in this demo)
)

// Constants Declaration
const (
	MAX_COUNT  = 1000          // Constant for maximum count
	NICK_NAME  = "Erren Yager" // Constant for nickname
	LET_PROCED = false         // Constant boolean value
)

// Skipping and resetting iota
const (
	X = iota // -
	Y        // 1
	_
	Z //3 (skipped index 2)
)

const (
	RESET = iota // 0 (new const block reset iota)
)

// Using iota for Bit Flags
const (
	READ    = 1 << iota // 1 << 0 = 0001 (1)
	WRITE               // 1 << 1 = 0010 (2)
	EXECUTE             // 1 << 2 = 0100 (4)
	DELETE              // 1 << 3 = 1000 (8)
)

// Dara Nasibi Enum using iota
type DaraNasibi int

const (
	VeryBad  DaraNasibi = iota // 0
	Bad                        // 1
	Neutral                    // 2
	Good                       // 3
	VeryGood                   // 4

)

// getDaraNasibi returns a human-readable string for the Dara Nasibi values.
func getDaraNasibi(dn DaraNasibi) string {
	switch dn {
	case VeryBad:
		return "Very Bad Luck 😞"
	case Bad:
		return "Bad Luck ☹️"
	case Neutral:
		return "Neutral Luck 😐"
	case Good:
		return "Good Luck 🙂"
	case VeryGood:
		return "Very Good Luck 😃"
	default:
		return "Unknown Luck 🤔"
	}
}

func main() {
	// ---------------------------
	// Basic I/O and Variable Examples
	// ---------------------------
	fmt.Println("Hello, World!")

	// Local variable 'Name' shadows the global variable 'Name'
	Name := "Abtin"

	// Demonstrate pass-by-value (does not change original value)
	passByValue(Name)
	fmt.Println("After passByValue, Name:", Name)

	// Demonstrate pass-by-reference (modifies original variable)
	passByReference(&Name)
	fmt.Println("After passByReference, Name:", Name)

	// Demonstrate control structures (if-else, switch, loops)
	demonstrateControlStructures()

	// Demonstrate various variable declarations and initializations
	variables()

	// Demonstrate constant usage
	const TEST string = "A"
	fmt.Println("The constant TEST is:", TEST)

	// ---------------------------
	// Arrays and Slices
	// ---------------------------
	// Arrays: fixed-size, indexed collections
	var a1 [10]int  // Array with default zero values
	a2 := [10]int{} // Same as above: explicitly zeroed
	fmt.Println("a1[7]:", a1[7])
	fmt.Println("a2[0]:", a2[0])

	// Slices: dynamic arrays
	var s1 []int          // Nil slice; no elements
	s1 = append(s1, 10)   // Append element (slice now has one element)
	s2 := make([]int, 10) // Slice with length 10, all zeroes
	s2 = append(s2, 11)   // Append an element (length becomes 11)
	fmt.Println("s1[0]:", s1[0])
	fmt.Printf("s1 length: %d, capacity: %d\n", len(s1), cap(s1))
	fmt.Println("s2[10]:", s2[10])
	fmt.Printf("s2 length: %d, capacity: %d\n", len(s2), cap(s2))

	// ---------------------------
	// Maps
	// ---------------------------
	// Maps: unordered key-value pairs
	m1 := make(map[string]string) // Empty map
	m2 := map[string]int{         // Map with initial values
		"Abtin":   22,
		"Fatemeh": 22,
	}
	var m3 map[int]bool // Nil map (uninitialized)

	fmt.Println("m2[\"Abtin\"]:", m2["Abtin"])
	fmt.Println("m2[\"Fatemeh\"]:", m2["Fatemeh"])

	// Assigning and checking map values
	m1["hello"] = "world"
	fmt.Printf("m1: %v\n", m1)
	fmt.Printf("m3: %v\n", m3)

	// Check for key existence in a map and delete an entry
	m := map[string]string{"Name": "Abtin"}
	if _, ok := m["Name"]; ok {
		println("Item exists")
	} else {
		println("Item does not exist")
	}

	if val, ok := m["Name"]; ok {
		println("Item exists:", val)
	} else {
		println("Item does not exist.")
	}

	delete(m, "Name")
	if _, ok := m["Name"]; ok {
		println("Item exists")
	} else {
		println("Item does not exist")
	}

	// ---------------------------
	// Structs, Methods, and Interfaces
	// ---------------------------
	// Create several Person instances using different initialization styles.
	person := Person{Name: "Abtin", Age: 22}
	fmt.Println(person.Introduce())

	p1 := Person{"Fatemeh", 21}
	fmt.Println(p1.Introduce())

	var p2 Person
	p2.Name = "Amir"
	p2.Age = 21
	fmt.Println(p2.Introduce())

	// Using new() returns a pointer initialized to zero values.
	p3 := new(Person)
	p3.Name = "Parsa"
	fmt.Println(p3.Introduce())

	// Modify Person's Name via pointer receiver method.
	p3.IntroducePointer("Dara")
	// Use a standalone function to show behavior (p3 is dereferenced)
	fmt.Println(IntroduceFunction(*p3))

	// Demonstrate optional argument using variadic parameters.
	fmt.Println(p3.IntroduceOptional())                  // Uses default name ("Dara")
	fmt.Println(p3.IntroduceOptional("Fatemeh", "Sara")) // Uses first optional argument ("Fatemeh")

	// Demonstrate setAge: setting an age less than 7 returns an error.
	if err := p3.setAge(1); err != nil {
		fmt.Println("setAge Error:", err)
	}
	if err := p3.setAge(21); err != nil {
		fmt.Println("setAge Error:", err)
	} else {
		fmt.Println("setAge succeeded. New Age:", p3.Age)
	}
	fmt.Println(p3.Introduce())

	// ---------------------------
	// Interfaces Demonstration
	// ---------------------------
	// The Speaker interface is defined below. Person now implements Speaker via its Speak() method.
	// We can assign a Person instance to a variable of type Speaker.
	p4 := Person{Name: "Abtin", Age: 22}
	var speaker Speaker = p4
	fmt.Println("Interface Speaker says:", speaker.Speak())

	// Comprehensive interface example with multiple implementations is shown in demoInterfaceExample().
	demoInterfaceExample()

	// ---------------------------
	// Pointers and Memory
	// ---------------------------
	pointerExample()

	// ---------------------------
	// Concurrency Example
	// ---------------------------
	go demoConcurrency()

	// ---------------------------
	// Error Handling Example
	// ---------------------------
	if err := demoErrorHandling(4); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Number is even")
	}

	// ---------------------------
	// Defer Example
	// ---------------------------
	demoDefer()

	// ---------------------------
	// Type Conversion Example
	// ---------------------------
	demoTypeConversion()

	// ---------------------------
	// Bitwise Shift Operations
	// ---------------------------
	leftShift()
	rightShift()

	// ---------------------------
	// Error
	// ---------------------------

	// The default value for Interface is nil, if we have nil then we do not have errors.
	// If we have, it means we have an error.
	err := ReturnSomeError()
	if err != nil {
		// do something with error
		fmt.Println("Error.")
	}
	// fmt.Println(ReturnSomeError())

	// ---------------------------
	// Inheritance in Go
	// ---------------------------
	demoStructEmbedding()

	// ---------------------------
	// Interface Casting in Go
	// ---------------------------

	// Create instances
	animal := Animal{Species: "Dog"}
	machine := Machine{Type: "Robot"}

	// Demonstrate interface casting
	fmt.Println("---- Casting Example with Animal ----")
	demoInterfaceCasting(animal)

	fmt.Println("---- Casting Example with Machine ----")
	demoInterfaceCasting(machine)

	// ---------------------------
	// Closure
	// ---------------------------
	demoClosure()

	// closure: function which returned by a function
	// closure inherits the parent function scope + its own scope

	// ---------------------------
	// Packages
	// ---------------------------

	// Using the greeting packages
	fmt.Println(greeting.SayHello("Abtin"))
	fmt.Println(greeting.SayGoodbye("Abtin"))

	// Using the mathutils package
	fmt.Println("Addition:", mathutils.Add(10, 5))
	fmt.Println("Multiplication:", mathutils.Multiply(4, 3))

	// Using alias package
	dateStr := alias.MyCustomString("2024-02-07")
	parsedDate, isValid := dateStr.IsDate()
	if isValid {
		fmt.Println("Valid date:", parsedDate.Format(time.RFC3339)) // Expected output: 2024-02-07T00:00:00Z
	} else {
		fmt.Println("Invalid date")
	}

	// **Using MyCustomString**
	dateStr = alias.MyCustomString("2025-03-10")
	parsedDate, isValid = dateStr.IsDate()
	fmt.Println("Is valid date?", isValid, "Parsed Date:", parsedDate)

	// Convert string to uppercase
	fmt.Println("Uppercase:", alias.MyCustomString("hello").ToUpper())

	// **Using MyCustomInt**
	num := alias.MyCustomInt(5)
	fmt.Println("Is Even:", num.IsEven())

	// compute factorial
	fact, err := num.Factorial()
	if err != nil {
		fmt.Println("Factorial Error:", err)
	} else {
		fmt.Println("Factorial:", fact)
	}

	// **Using MyOperation (Function Alias)**
	result1 := alias.ApplyOperation(10, 5, alias.Add)
	result2 := alias.ApplyOperation(10, 5, alias.Multiply)

	fmt.Println("Addition Result:", result1)
	fmt.Println("Multiplication Result:", result2)

	// Invalid date string
	invalidStr := alias.MyCustomString("Hello World")
	parsedDate, isValid = invalidStr.IsDate()
	if isValid {
		fmt.Println("Valid date:", parsedDate.Format(time.RFC3339))
	} else {
		fmt.Println("Invalid date") // Expected output: Invalid date
	}

	// Using stringutils package
	str := " Hello, GoLang!"

	// Convert to Upper and Lower case
	fmt.Println("UpperCase:", stringutils.ToUpperCase(str))
	fmt.Println("LowerCase:", stringutils.ToLowerCase(str))

	// Check if a substring exists
	fmt.Println("Contains 'Go':", stringutils.Contains(str, "Go"))

	// Trim leading and trailing spaces
	fmt.Println("Trimmed:", stringutils.TrimSpaces(str))

	// Replace substring
	fmt.Println("Replaced 'GoLang' with 'Gophers':", stringutils.ReplaceSubstring(str, "GoLang", "Gophers"))

	// Split string
	splitStr := stringutils.SplitString("Go,Python,Java", ",")
	fmt.Println("Split:", splitStr)

	// Using datastructures package
	// ----- Linked List Example -----
	ll := datastructures.LinkedList{}
	ll.Append("first")
	ll.Append("second")
	ll.Prepend("zero")
	ll.Print() // Expected: zero -> first -> second -> nil

	// ----- Stack Example -----
	var s datastructures.Stack
	s.Push(10)
	s.Push(20)
	top, _ := s.Peek()
	fmt.Println("Stack top:", top) // Expected: 20
	item, _ := s.Pop()
	fmt.Println("Popped:", item) // Expected: 20

	// ----- Queue Example -----
	var q datastructures.Queue
	q.Enqueue(100)
	q.Enqueue(200)
	front, _ := q.Peek()
	fmt.Println("Queue front:", front) // Expected: 100
	item, _ = q.Dequeue()
	fmt.Println("Dequeued:", item) // Expected: 100

	// ----- Binary Search Tree Example -----
	bst := datastructures.BinarySearchTree{}
	bst.Insert(50)
	bst.Insert(30)
	bst.Insert(70)
	bst.Insert(20)
	bst.Insert(40)
	fmt.Println("BST InOrder:", bst.InOrderTraversal())
	fmt.Println("Search 40 in BST:", bst.Search(40))
	fmt.Println("Search 90 in BST:", bst.Search(90))

	// ----- Priority Queue Example -----
	items := []*datastructures.PriorityQueueItem{
		{Value: "Task1", Priority: 3},
		{Value: "Task2", Priority: 1},
		{Value: "Task3", Priority: 2},
	}
	pq := make(datastructures.PriorityQueue, len(items))
	for i, item := range items {
		pq[i] = item
		pq[i].Index = i
	}
	heap.Init(&pq)
	heap.Push(&pq, &datastructures.PriorityQueueItem{Value: "Task4", Priority: 0})
	itemPopped := heap.Pop(&pq).(*datastructures.PriorityQueueItem)
	fmt.Println("Highest Priority Task:", itemPopped.Value) // Expected: Task4

	// ----- Binary Search Algorithm Example -----
	sortedArr := []int{1, 3, 5, 7, 9}
	index := datastructures.BinarySearch(sortedArr, 7)
	fmt.Println("Index of 7:", index) // Expected: 3

	// ----- Standard Package Demonstrations -----
	datastructures.DemoContainerList()
	datastructures.DemoContainerHeap()
	datastructures.DemoContainerRing()
	datastructures.DemoSort()
	datastructures.DemoMathRand()

	// exampleFunction is a simple function that prints a message.
	fmt.Println("=== Delayed Function Execution Demo ===")

	// Wrap exampleFunction with delay
	delayedFunction := delay(exampleFunction)

	// Call the delayed function
	delayedFunction()

	// Wrap add function with logExecutionDecorator
	decoratedAdd := logExecutionDecorator(func(i1, i2 int) int { return i1 + i2 })

	// Call the decorated function
	decoratedAdd(5, 10)

	// ---------------------------
	// Iota
	// ---------------------------

	// Basic iota
	demoBasicIota()

	// Custom step values
	demoBasicIota2()

	// Skipping and Resetting iota
	demoBasicIota3()

	// Using iota for Bit FLAGS
	demoBasicIota4()

	// Using iota in Structs and Switch Statements
	demoBasicIota5()

	demoBasicIota6()

	// ---------------------------
	// Generics
	// ---------------------------

	demoGeneric1()

	demoGeneric2()

	demoGeneric3()

	demoGeneric4()

	demoGeneric5()

	demoGeneric6()

	demoGeneric7()

	// ---------------------------
	// Empty Interface{}
	// ---------------------------

	// Since Go 1.18, generics (T any) are preferred over interface{} for type safety.
	demoEmptyInterface1()
	demoEmptyInterface2()
	demoEmptyInterface3()
	demoEmptyInterface4()
	demoEmptyInterface5()
	demoEmptyInterface6()

	// ---------------------------
	// Map
	// ---------------------------
	demoMap()

	// ------------------------------------------------------
	// Ways to get input from terminal in Go
	// ------------------------------------------------------

	// fmt.Scan, which reads space-separated values.
	var name string
	fmt.Print("Enter your name:")
	fmt.Scan(&name) // Reads input into `name` (stops at space/newline)
	fmt.Println("Hello, ", name)

	// fmt.Scanln, reads an entire line (including space), use fmt.Scanln
	fmt.Println(&name) //  Reads until new line
	fmt.Println("Hello", name)

	// bufio.NewReader (Full Line Input)
	// for full-line input, we use bufio.NewReader(os.Stdin)

	reader := bufio.NewReader(os.Stdin) // Create a buffered reader
	name1, _ := reader.ReadString('\n') // Reads input until newline
	fmt.Println("Hello", name1)

	// os.Stdin.Read, which reads raw bytes slice
	fmt.Print("Enter something:")
	input := make([]byte, 100) // Allocate 100
	n, _ := os.Stdin.Read(input)
	fmt.Println("You entered:", string(input[:n]))
}

// comparable --> only types that support `==` and `!=` comparisons.

// FindIndex searches for an item in a slice and returns its index.
func FindIndex2[T comparable](arr []T, target T) int {
	for i, v := range arr {
		if v == target {
			return i
		}
	}

	return -1
}

// Define Custom constraint with interfaces

// Numeric is a custom constraint that allows only int and float types.
type Numeric interface {
	int | float64
}

// Sum calculates the sum of a slice of numeric values.
func Sum[T Numeric](numbers []T) T {
	var total T

	for _, num := range numbers {
		total += num
	}

	return total
}

// Using ~ for type Aliases

// MyInt is a custom type alias of int.
type MyInt int

// Numeric is constraint that allows int and all aliases of int.
type Numeric2 interface {
	~int | ~float64
}

// Multiply multiplies each element in a slice by a factor.
func Multiply[T Numeric](nums []T, factor T) []T {
	result := make([]T, len(nums))
	for i, v := range nums {
		result[i] = v * factor
	}

	return result
}

// Combining Multiple constraints

// StringOrInt allows only int or string values.
type StringOrInt interface {
	int | string
}

// PrintValues prints values of allowed types.
func PrintValues[T StringOrInt](values []T) {
	for _, v := range values {
		fmt.Println("Values", v)
	}
}

/*
	Map
*/
// transfer each element of T to U with given mapper function.

// Generic map function: Transforms a slice of T into a slice of U.
func Map[T, U any](s []T, f func(T) U) []U {
	result := make([]U, len(s))
	for i, v := range s {
		result[i] = f(v) // Apply function f to each element
	}
	return result
}

// Example: convert a slice of integers to strings.
func intToString(n int) string {
	return fmt.Sprintf("Number: %d", n)
}

// Example: convert a slice of strings to uppercase
func ToUpperCase(s string) string {
	return stringutils.ToUpperCase(s)
}

func demoMap() {
	fmt.Println("=== Generic Map Function Demo ===")

	// Example 1: Map ints to string
	numbers := []int{1, 2, 3, 4, 5}
	stringsResult := Map(numbers, intToString)
	fmt.Println("Mapped Integers to Strings:", stringsResult)

	// Example 2: Map strings to uppercase
	words := []string{"hello", "world", "golang"}
	upperWords := Map(words, ToUpperCase)
	fmt.Println("Mapped strings to Uppercase:", upperWords)
}

/*
	Empty Interface
*/

// storing Any Type in an Empty Interface
func demoEmptyInterface1() {
	var X interface{} // Can hold any type

	X = 42         // Assign an int
	fmt.Println(X) // Output: 42

	X = "Hello"    // Assign a string
	fmt.Println(X) // Output: Hello

	X = 3.14       // Assign a float
	fmt.Println(X) // Output: 3.14
}

// Functions That Accepts Any Type
func PrintAny(value interface{}) {
	fmt.Println("Value received:", value)
}

func demoEmptyInterface2() {
	PrintAny(100)         // int
	PrintAny("GoLang")    // string
	PrintAny(3.14)        // float
	PrintAny([]int{1, 2}) //slice
}

// Type Assertion: Extracting the Actual Type
func checkType(value interface{}) {
	// Type ASSERTION
	str, ok := value.(string)

	if ok {
		fmt.Println("This is a string", str)
	} else {
		fmt.Println("Not a string")
	}
}

func demoEmptyInterface3() {
	checkType("Hello") // This is a string: Hello
	checkType(55)      // NOt a string
}

// Using a Type Switch for Multiple Types
func detectType(value interface{}) {
	switch v := value.(type) {
	case int:
		fmt.Println("Integer:", v)
	case string:
		fmt.Println("String:", v)
	case float64:
		fmt.Println("Float:", v)
	default:
		fmt.Println("Unknown Type!")
	}
}

func demoEmptyInterface4() {
	detectType(19)          // Integer: 19
	detectType("GOlang")    // String: GoLang
	detectType(3.14)        // Float: 3.14
	detectType([]int{1, 2}) // Unknown Type!

}

// Storing Multiple Types in a Slice
func demoEmptyInterface5() {
	data := []interface{}{42, "Hello", 3.14, true}

	for _, v := range data {
		fmt.Println("Value:", v)
	}
}

// Using interface{} in JSON Unmarshaling

// JSON data as a string
var jsonData = `{"name": "Alice", "age": 25, "isStudent": false}`

func demoEmptyInterface6() {
	var result map[string]interface{} //  Dynamic map

	// Unmarshal JSON into map
	err := json.Unmarshal([]byte(jsonData), &result)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Decoded JSON:", result)

	// Extracting values safely
	name := result["name"].(string)
	age := int(result["age"].(float64)) // JSON numbers default to float64
	isStudent := result["isStudent"].(bool)

	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Is Student:", isStudent)
}

// Implementing a Generic Functions with interface{} (Before Go 1.18 Generics)
func printAll(values ...interface{}) {
	for _, v := range values {
		fmt.Println("Value:", v)
	}
}

func demoEmptyInterface7() {
	printAll(10, "Hello", 3.14, true)
}

/*
	Generics
*/

// Basic Generics Example (Type Parameters)

// Generic function that accepts any types and returns it.
func Identity[T any](value T) T {
	return value
}

func demoGeneric1() {
	fmt.Println("Generic Identity Function:")
	fmt.Println(Identity(43))       // Works with int
	fmt.Println(Identity("Golang")) // Works with string
	fmt.Println(Identity(3.15))     // Works with float
}

// Generic with Multiple Type Parameters
// Function that swaps two values of different types
func Swap[T, U any](a T, b U) (U, T) {
	return b, a
}

func demoGeneric2() {
	fmt.Println("Generic Swap Function:")
	a, b := Swap(100, "Hello")
	fmt.Println(a, b) // output: Hello 100
}

// Generics Structs
type Box[T any] struct {
	Value T
}

func demoGeneric3() {
	fmt.Println("Generic Struct Example:")

	intBox := Box[int]{Value: 42}
	strBox := Box[string]{Value: "GoLang"}

	fmt.Println("Int Box:", intBox.Value)    //42
	fmt.Println("String Box:", strBox.Value) //GoLang
}

// Generic Methods (Functions on Generic Structs)
// Generic Stack implementation

type Stack[T any] struct {
	items []T
}

// Push adds an element to stack
func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

// Pop removes and returns the top element
func (s *Stack[T]) Pop() (T, bool) {
	if len(s.items) == 0 {
		var zero T // Default zero Value
		return zero, false
	}

	last := len(s.items) - 1
	item := s.items[last]
	s.items = s.items[:last]
	return item, true
}

func demoGeneric4() {
	fmt.Println("Generic Stack Example:")

	stack := Stack[int]{}
	stack.Push(10)
	stack.Push(20)

	value, _ := stack.Pop()
	fmt.Println("Popped", value) // 20
}

// Generic Constraints
// (Comparable, ~T)
// Constraints restrict which types can be used.
// comparable allows equality comparisons.

// FindIndex returns the index of an item in a slice.
func FindIndex[T comparable](arr []T, item T) int {
	for i, v := range arr {
		if v == item {
			return i
		}
	}

	return -1
}

func demoGeneric5() {
	fmt.Println("Generics with constraints example:")

	nums := []int{10, 20, 30, 40}
	fmt.Println("Index of 30:", FindIndex(nums, 30)) // 2

	words := []string{"Go", "Python", "Java"}
	fmt.Println("Index of Python:", FindIndex(words, "Python"))
}

// Generic Interface

// Define a generic Interface
type Summable[T any] interface {
	Sum() T
}

// Implement the interface for an IntList
type IntList []int

func (l IntList) Sum() int {
	total := 0
	for _, num := range l {
		total += num
	}
	return total
}

func demoGeneric6() {
	fmt.Println("Generic INterface Example:")

	list := IntList{1, 2, 3, 4, 5}
	fmt.Println("Sum:", list.Sum()) // 15
}

// Generic Maps

// GEneric Maps WRAPPER
type Dictionary[K comparable, V any] struct {
	data map[K]V
}

// NewDictionary create a new dictionary.
func NewDictionary[K comparable, V any]() *Dictionary[K, V] {
	return &Dictionary[K, V]{data: make(map[K]V)}
}

// Set adds a kwy-value pair
func (d *Dictionary[K, V]) Set(key K, value V) {
	d.data[key] = value
}

// Get retrieves a value by key
func (d *Dictionary[K, V]) Get(key K) (V, bool) {
	val, exists := d.data[key]
	return val, exists
}

func demoGeneric7() {
	fmt.Println("Generic Dictionary Example:")

	dict := NewDictionary[string, int]()
	dict.Set("Alice", 30)
	dict.Set("Bov", 25)

	age, _ := dict.Get("Alice")
	fmt.Println("Alices's Age:", age) // 25
}

/*
	Basic iota examples
*/

func demoBasicIota() {
	fmt.Println("Basic iota example: ")
	fmt.Println("A =", A)
	fmt.Println("B =", B)
	fmt.Println("C =", C)
	fmt.Println("D =", D)
}

func demoBasicIota2() {
	fmt.Println("Memory Sizes Using iota:")
	fmt.Println("KB =", KB)
	fmt.Println("MB =", MB)
	fmt.Println("GB =", GB)
	fmt.Println("TB =", TB)
}

func demoBasicIota3() {
	fmt.Println("Skipping and Resetting iota:")
	fmt.Println("X =", X)
	fmt.Println("Y =", Y)
	fmt.Println("Z =", Z)
	fmt.Println("RESET =", RESET)
}

func demoBasicIota4() {
	fmt.Println("Bit Flags Using iota: ")
	fmt.Println("READ =", READ)
	fmt.Println("WRITE =", WRITE)
	fmt.Println("EXECUTE =", EXECUTE)
	fmt.Println("DELETE =", DELETE)
}

func logMessage(level LogLevel) {
	switch level {
	case DEBUG:
		fmt.Println("Debug Message")
	case INFO:
		fmt.Println("Info Message")
	case WARN:
		fmt.Println("Warning Message")
	case ERROR:
		fmt.Println("Error Message")
	default:
		fmt.Println("Unknown Log Message")
	}
}

func demoBasicIota5() {
	fmt.Println("Using iota with Switch")
	logMessage(DEBUG) // DEBUG Message
	logMessage(ERROR) // Error Message
	logMessage(5)     // Unknown Log Level
}

func demoBasicIota6() {
	fmt.Println("=== Dara Nasibi Enum Example ===")

	// Print all Dara Nasibi levels
	fmt.Println("VeryBad:", getDaraNasibi(VeryBad))
	fmt.Println("Bad:", getDaraNasibi(Bad))
	fmt.Println("Neutral:", getDaraNasibi(Neutral))
	fmt.Println("Good:", getDaraNasibi(Good))
	fmt.Println("VeryGood:", getDaraNasibi(VeryGood))

	// Example of handling an unknown case
	var unknownValue DaraNasibi = 10
	fmt.Println("Unknown:", getDaraNasibi(unknownValue))
}

/*
	Basic Functions Demonstrating Parameter Passing
*/

// passByValue shows that passing a string by value does not modify the original variable.
func passByValue(s string) {
	s = "felan" // Modifies only the local copy.
}

// passByReference demonstrates that passing a pointer allows modification of the original variable.
func passByReference(s *string) {
	*s = "felan" // Dereferences and modifies the original value.
}

// variables demonstrates different variable declarations and initializations.
func variables() {
	var var1 int32 = 32
	fmt.Println("var1 (int32):", var1)

	var2 := "string"
	fmt.Println("var2 (string):", var2)

	var3 := 32.2 // Defaults to float64.
	fmt.Println("var3 (float64):", var3)

	v1, v2, v3 := false, "hello", 23
	fmt.Println("v1 (bool), v2 (string), v3 (int):", v1, v2, v3)
}

/*
Control Structures Examples
*/
func demonstrateControlStructures() {
	// if-else with an initialization statement.
	if x := 10; x > 5 {
		fmt.Println("x is greater than 5")
	} else {
		fmt.Println("x is less than or equal to 5")
	}

	// switch statement example.
	switch day := 3; day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	default:
		fmt.Println("Unknown day")
	}

	// For loop example.
	for i := 0; i < 5; i++ {
		fmt.Println("For loop iteration:", i)
	}

	// Infinite loop (commented out to prevent blocking execution)
	// for {
	//     fmt.Println("This loop runs forever")
	// }

	s := []int{1, 2, 3, 4, 5}

	for i, v := range s {
		fmt.Println("index: ", i, "value: ", v)
	}

	m := map[string]int{
		"Abtin":   22,
		"Tara":    23,
		"Fatemeh": 20,
	}

	for m, v := range m {
		fmt.Println("key :", m, "value :", v)
	}

	for i := range 5 {
		fmt.Println("num: ", i)
	}

	for i := 0; i < 4; i++ {
		switch i {
		case 0:
			fmt.Println("F Y")
		case 1:
			fmt.Println("U O")
		case 2:
			fmt.Println("C U")
		case 3:
			fmt.Println("k")
		}

	}

}

/*
	Structs, Methods, and Interfaces
*/

// Person struct represents a person with a Name and Age.
type Person struct {
	Name string
	Age  int
}

// Constructor example in Go
func NewPerson(fn, ln string, age int) Person {
	return Person{
		Name: fn + " " + ln,
		Age:  age,
	}
}

func NewPersonPointer(fn, ln string, age int) *Person {
	return &Person{
		Name: fn + " " + ln,
		Age:  age,
	}
}

// Introduce is a method with a value receiver that returns a formatted introduction.
// Since it uses a value receiver, modifications inside do not affect the original instance.
func (p Person) Introduce() string {
	return fmt.Sprintf("Hi, I'm %s, %d years old.", p.Name, p.Age)
}

// IntroducePointer is a method with a pointer receiver that modifies the Person's Name.
func (p *Person) IntroducePointer(newName string) {
	p.Name = newName // Modifies the original Person instance.
}

// IntroduceFunction is a standalone function that takes a Person (by value) and returns an introduction string.
func IntroduceFunction(p Person) string {
	return fmt.Sprintf("Hi, I'm %s, %d years old.", p.Name, p.Age)
}

// IntroduceOptional uses variadic parameters to allow an optional new name.
// If a name is provided, it updates the Person's Name; otherwise, it uses the existing Name.
func (p *Person) IntroduceOptional(names ...string) string {
	if len(names) > 0 {
		p.Name = names[0] // Use the first provided name.
	}
	return fmt.Sprintf("Hi, I'm %s, %d years old.", p.Name, p.Age)
}

// setAge sets the Person's Age if the provided age is valid (>= 7).
// If the age is less than 7, it returns an error.
func (p *Person) setAge(age int) error {
	if age < 7 {
		return fmt.Errorf("invalid age: %d (must be at least 7 years old)", age)
	}
	p.Age = age
	return nil
}

/*
	Interfaces in Go
	----------------
	An interface defines a set of method signatures. Any type that implements these methods
	automatically satisfies the interface (implicit implementation).
*/

// Implicit: Happens automatically, without requiring manual specification.
/*
	// Explicit Declaration (Type Specified)
	var num int = 10

	// Implicit Type Inference (Using `:=` Operator)
	num2 := 20
*/

// Error is a built-in function in Go
/*
	type error interface {
		Error() string
	}
*/

// Below code use the package of errors
/*
	func ReturnSomeError() error {
		return errors.New("This is an error.")
	}
*/

func ReturnSomeError() error {
	return myCustomError{
		code: 30,
		msg:  "system panic :()",
	}
}

// Custom Error
type myCustomError struct {
	code int
	msg  string
}

func (e myCustomError) Error() string {
	return fmt.Sprintf("%d - %s", e.code, e.msg)
}

// Speaker interface defines behavior for types that can "speak".
type Speaker interface {
	Speak() string
}

// Speak is a method on Person that satisfies the Speaker interface.
// It returns a greeting message.
func (p Person) Speak() string {
	return fmt.Sprintf("Hello, my name is %s.", p.Name)
}

// demoInterfaceExample demonstrates the use of interfaces by assigning a Person to a Speaker variable.
func demoInterfaceExample() {
	// Create a Person instance.
	p := Person{Name: "Abtin", Age: 22}

	// Since Person implements Speak(), it satisfies the Speaker interface.
	var speaker Speaker = p
	fmt.Println("demoInterfaceExample - Speaker says:", speaker.Speak())

	// --- Additional Interface Concepts ---
	// You can use type assertions or type switches with interfaces.
	describeInterface(speaker)
}

// describeInterface demonstrates type assertion on an interface.
func describeInterface(i interface{}) {
	// Try to assert i as a Person.
	if person, ok := i.(Person); ok {
		fmt.Println("describeInterface: It is a Person with Name:", person.Name)
	} else {
		fmt.Println("describeInterface: Unknown type")
	}
}

/*
Concurrency: Goroutines and Channels
*/
func demoConcurrency() {
	ch := make(chan string)
	go func() {
		ch <- "Hello from Goroutine!"
	}()
	message := <-ch
	fmt.Println("demoConcurrency:", message)
}

/*
Error Handling
--------------
Functions in Go can return multiple values, and error handling is done explicitly.
*/
func demoErrorHandling(n int) error {
	if n%2 != 0 {
		return fmt.Errorf("number is odd")
	}
	return nil
}

/*
Defer Statement
---------------
'defer' schedules a function call to be run after the surrounding function returns.
*/
func demoDefer() {
	defer fmt.Println("demoDefer: This will be printed last!")
	fmt.Println("demoDefer: This will be printed first!")
}

/*
Type Conversion
---------------
Go allows explicit conversion between types.
*/
func demoTypeConversion() {
	var intVar int = 42
	var floatVar float64 = float64(intVar)
	fmt.Println("Converted int to float:", floatVar)

	var strVar string = fmt.Sprintf("%d", intVar)
	fmt.Println("Converted int to string:", strVar)
}

/*
Bitwise Shift Operations
--------------------------
Left shift (<<) multiplies a number by 2^n, while right shift (>>) divides (rounding down).
*/
func leftShift() {
	var a int = 4       // Binary: 0100
	var result = a << 2 // Shifting left by 2 (0100 becomes 10000, which is 16)
	fmt.Printf("leftShift: a: %b (binary), %d (decimal)\n", a, a)
	fmt.Printf("leftShift: a << 2: %b (binary), %d (decimal)\n", result, result)
}

func rightShift() {
	var a int = 16      // Binary: 10000
	var result = a >> 2 // Shifting right by 2 (10000 becomes 100, which is 4)
	fmt.Printf("rightShift: a: %b (binary), %d (decimal)\n", a, a)
	fmt.Printf("rightShift: a >> 2: %b (binary), %d (decimal)\n", result, result)
}

/*
Pointers and Memory in Go
--------------------------
Pointers store the memory address of a variable. Go pointers do not support arithmetic.
Memory is managed on the stack (for local variables) or on the heap (for dynamically allocated data).
*/
func pointerExample() {
	// 'a' is a primitive variable stored on the stack.
	var a int = 100
	var ptr *int // Declare a pointer to int.
	ptr = &a     // Assign the address of 'a' to ptr.

	// Dereferencing the pointer to access the value of 'a'.
	fmt.Println("pointerExample: Address of 'a':", ptr)
	fmt.Println("pointerExample: Value of 'a' through ptr:", *ptr)

	// Demonstrate modifying a struct field via a pointer.
	p := &Person{Name: "Parsa", Age: 21}
	changeName(p)
	fmt.Println("pointerExample: Updated Name (after changeName):", p.Name)

	// Safe pointer usage: initialize a uint variable and assign its address.
	var x *uint // x is a nil pointer initially.
	var y uint  // Regular uint variable.
	y = 10
	x = &y // x now points to y.
	fmt.Println("pointerExample: x points to value:", *x)

	// Demonstrate pointer-to-pointer (although rarely needed in Go).
	z := &y
	fmt.Println("pointerExample: Address of y through z:", z)
	fmt.Println("pointerExample: Value of y through z:", *z)
}

// changeName modifies the Name field of a Person via a pointer.
func changeName(p *Person) {
	p.Name = "Batman"
}

/*
	Struct Embedding (Inheritance-like Behavior in Go)
	--------------------------------------------------
	Go does not have classical inheritance, but struct embedding allows us to reuse code.
	Instead of subclassing, we embed one struct inside another to extend functionality.
*/

// Base struct (similar to a parent class)
type PersonBase struct {
	Name string
	Age  int
}

// Method for PersonBase
func (p PersonBase) Introduce() string {
	return fmt.Sprintf("Hi, I'm %s, and I'm %d years old.", p.Name, p.Age)
}

// Derived struct (similar to a subclass)
// Embedding `PersonBase` inside `IranianPersonStruct`
// This allows `IranianPersonStruct` to inherit the fields and methods of `PersonBase`
type IranianPersonStruct struct {
	PersonBase // Embedded struct (inherits Name and Age fields)
	Country    string
}

// SpeakPersian method for IranianPersonStruct (new behavior for this "subclass").
func (i IranianPersonStruct) SpeakPersian() string {
	return fmt.Sprintf("%s says: سلام! (Hello in Persian)", i.Name)
}

// Another Derived Struct (AsianPersonStruct)
type AsianPersonStruct struct {
	PersonBase // Embedded struct (inherits Name and Age fields)
	Country    string
}

// EatWithChopsticks method for AsianPersonStruct
func (a AsianPersonStruct) EatWithChopsticks() string {
	return fmt.Sprintf("%s from %s is eating with chopsticks.", a.Name, a.Country)
}

// Function to demonstrate Struct Embedding (Inheritance-like Behavior)
func demoStructEmbedding() {
	// Create an instance of IranianPersonStruct
	iranian := IranianPersonStruct{
		PersonBase: PersonBase{Name: "Ali", Age: 30}, // Inheriting Name and Age
		Country:    "Iran",
	}

	// Create an instance of AsianPersonStruct
	asian := AsianPersonStruct{
		PersonBase: PersonBase{Name: "Lee", Age: 25}, // Inheriting Name and Age
		Country:    "China",
	}

	// Call methods inherited from PersonBase
	fmt.Println(iranian.Introduce()) // Calls Introduce() from PersonBase
	fmt.Println(asian.Introduce())   // Calls Introduce() from PersonBase

	// Call subclass-specific methods
	fmt.Println(iranian.SpeakPersian())    // Calls SpeakPersian()
	fmt.Println(asian.EatWithChopsticks()) // Calls EatWithChopsticks()
}

// Interface Casting

// Define an interface
type Communicator interface {
	Communicate() string
}

// Define two different struct types implementing Communicator
type Animal struct {
	Species string
}

type Machine struct {
	Type string
}

// Implement the Communicate() method for Animal
func (a Animal) Communicate() string {
	return fmt.Sprintf("The %s make a sound.", a.Species)
}

// Implement the Communicate() method for Machine
func (m Machine) Communicate() string {
	return fmt.Sprintf("The %s machine beeps!", m.Type)
}

// Function to demonstrate interface casting (type assertion and type switch)
func demoInterfaceCasting(c Communicator) {
	//1. **Type Assertion**
	animal, ok := c.(Animal)
	if ok {
		fmt.Println("Type Assertion: communicator is an Animal: ", animal.Species)
	} else {
		fmt.Println("Type Assertion: Communicator is NOT an Animal")
	}

	// Machine Type Assertion
	machine, ok := c.(Machine)
	if ok {
		fmt.Println("Type Assertion: Communicator is a Machine: ", machine.Type)
	} else {
		fmt.Println("Type Assertion: Communicator is NOT a Machine")
	}

	// 2. **Type Switch (More Flexible)**
	switch v := c.(type) {
	case Animal:
		fmt.Println("Type Switch: communicator is an Animal -", v.Communicate())
	case Machine:
		fmt.Println("Type Switch: Communicator is a Machine -", v.Communicate())
	default:
		fmt.Println("Type Switch: Unknown Type")
	}
}

// makeCounter returns a closure that keeps track of the count value.
func makeCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

// makeMultiplier returns a closure that multiplies its input by a fixed factor.
func makeMultiplier(factor int) func(int) int {
	return func(n int) int {
		return n * factor
	}
}

// makeCacheChecker returns a closure that caches function results to avoid redundant calculations.
func makeCacheChecker(fn func(int) int) func(int) int {
	cache := make(map[int]int) // Cache storage

	return func(n int) int {
		if result, found := cache[n]; found {
			fmt.Println("Cache hit for", n) // Found in cache
			return result
		}
		fmt.Println("Cache miss for", n) // Not in cache, computing
		result := fn(n)                  // Compute result
		cache[n] = result                // Store result in cache
		return result
	}
}

// getSequence returns a closure that generates a sequence of numbers.
// Every call to this closure returns a the next number in sequence.
func getSequence(start int) func() int {
	current := start
	return func() int {
		result := current
		current++
		return result
	}
}

// demoClosure demonstrate closures by using a counter, a multiplier function, a cache checker, and a sequence generator.
func demoClosure() {
	fmt.Println("=== closure Demo ===")

	// Example: Counter Closure
	counter := makeCounter()
	fmt.Println("Counter:", counter()) // 1
	fmt.Println("Counter:", counter()) // 2
	fmt.Println("Counter:", counter()) // 3

	// Example: Multiplier Closure
	double := makeMultiplier(2)
	triple := makeMultiplier(3)
	fmt.Println("Double of 4:", double(4)) // 8
	fmt.Println("Double of 4:", double(5)) // 10
	fmt.Println("Triple of 4:", triple(4)) // 12

	// Example: Cache Checker
	fmt.Println("=== Cache Checker Demo ===")
	multiplierWithCache := makeCacheChecker(double)

	fmt.Println("Compute 4 x 2:", multiplierWithCache(4))       // Cache miss, computes 8
	fmt.Println("Compute 5 x 2:", multiplierWithCache(5))       // Cache miss, computes 10
	fmt.Println("Compute 4 x 2 again:", multiplierWithCache(4)) // Cache hit, returns 8
	fmt.Println("Compute 5 x 2 again:", multiplierWithCache(5)) // Cache hit, returns 10

	// Example: getSequence Closure
	fmt.Println("==== getSequence Demo ===")
	sequence1 := getSequence(100) // Starts from 100
	sequence2 := getSequence(200) // Starts from 200

	fmt.Println("Sequence1:", sequence1()) // 100
	fmt.Println("Sequence1:", sequence1()) // 101
	fmt.Println("Sequence1:", sequence1()) // 102

	fmt.Println("Sequence2:", sequence2()) // 200
	fmt.Println("Sequence2:", sequence2()) // 201
	fmt.Println("Sequence2:", sequence2()) // 202

}

// delay wraps a function and adds a delay before executing it.
func delay(f func()) func() {
	return func() {
		fmt.Println("Starting ...")
		time.Sleep(time.Second * 2) // Introduce a delay
		f()                         // Execute the actual function
		fmt.Println("Done")
	}
}

// logExecutionDecorator decorates a function with arguments and a return value.
func logExecutionDecorator(f func(int, int) int) func(int, int) int {
	return func(a, b int) int {
		fmt.Printf("Executing function with inputs: %d, %d\n", a, b)
		result := f(a, b)
		fmt.Printf("Function result: %d\n", result)
		return result
	}
}

// exampleFunction is a simple function that print a message
func exampleFunction() {
	fmt.Println("Executing the actual function!")
}

/*
func checkSystemArchWithShift() {
	// Get the size of a pointer to an integer
	ptrSize := unsafe.sizeof(*new(int))

	// Left shift 1 by the number of bits to check if we are dealing with a 64-bit or 32-bit machine.
	// This is a simple demonstration that works because a 32-bit pointer is represented by 4 bytes (32 bits),
	// and a 64-bit pointer by 8 bytes (64 bits).
	shiftedPtrSize := ptrSize << 3 // Shift left by 3 bits (equivalent to multiplying by 8) to get the bit size.

	// Check if it's 32-bit or 64-bit
	if shiftedPtrSize == 32 {
		fmt.Println("This is a 32-bit architecture.")
	} else if shiftedPtrSize == 64 {
		fmt.Println("This is a 64-bit architecture.")
	} else {
		fmt.Println("Unknown architecture size:", shiftedPtrSize)
	}
}
*/

// Notes for Go language learners:
/*
Go data types:
- Primitive types:
  - Integers: uint, int, uint8, int8, etc.
  - Floating-point: float32, float64
  - Boolean: bool
  - String: string
  - Byte: 8-bit data (alias for uint8)
  - Rune: 32-bit Unicode character
  - Complex numbers: complex64, complex128

- Composite types:
  - Array: Fixed-size list
  - Slice: Dynamic array (resizable)
  - Map: Key-value pairs (like a hash-map)
  - Struct: A custom data type to hold different types of data
  - Interface: Defines methods for types to implement
  - Channel: Used for communication between goroutines (concurrency)

Zero values:
- Every type has a default "zero" value if not explicitly initialized:
  - int: 0
  - string: ""
  - bool: false
  - pointer: nil
  - numbers: 0
  - arrays: zero values for each element
  - slices, maps, channels: nil

Variable declaration:
- With explicit type and zero value: var VARIABLE <TYPE>
- With initialization + type inference: VARIABLE := <VALUE>
  - Numbers default to int.
  - Floats default to float64.
  - Strings default to string.
- Multiple variables can be declared together using a var block:
  - var (v1 TYPE1, v2 TYPE2, ...)

Constants:
- Declare constants with `const`:
  - const VARIABLE = VALUE
  - const (
      C1 TYPE1 = VALUE
      C2 TYPE2 = VALUE
    )

Naming conventions:
- Exported names (accessible outside the package) use PascalCase.
- Unexported names (accessible only within the package) use camelCase.
- Constants are often written in uppercase (C_STYLE_CONSTANT).

Pointers:
- ptr --> address --> 64 bit or 8 byte (arch is 64 bit)
- c --> ptr arithmetic
- stack --> var = primpitive (address value in heap and you can not change this address)
- address are the addresss in the stack

Using `make` function:
- `make` is used to initialize maps, slices, and channels in Go.

Defer, Panic, Recover:
- `defer` is used to schedule a function to be executed after the surrounding function finishes.
- `panic` is used to terminate the program immediately (usually in error cases).
- `recover` is used to regain control of a panicking program.

Type Conversion:
- Go allows you to explicitly convert between types using the type name (e.g., `float64(10)`).
*/

// Pointers and Memory in Go

/*
Pointers in Go:
- A pointer is a variable that stores the memory address of another variable.
- In Go, pointers do not allow arithmetic like in C, which means you cannot modify the memory address stored in the pointer directly (i.e., no pointer arithmetic).
- Pointers are used for referencing and modifying variables indirectly.

	Example of pointer declaration and usage:
	  var ptr *int   // A pointer to an integer
	  var a int = 10
	  ptr = &a       // The address of 'a' is assigned to ptr

	- &a gives the memory address of the variable 'a' (address of 'a').
	- *ptr dereferences the pointer to access the value at the address it is pointing to.

Memory Allocation in Go (Stack and Heap):
- Go uses stack and heap for memory allocation.
  - **Stack**: The stack holds local variables (primitive types, references, etc.), and variables are allocated in the stack at compile-time. The stack is fast and has limited size.
  - **Heap**: The heap holds variables that are dynamically allocated during runtime. These are typically large or long-lived objects (e.g., slices, maps, structs).

Pointers in the Stack:
- When you declare a variable of a primitive type (e.g., `int`), the value is stored in the stack, and its address is accessible via a pointer.
- The address of the variable is stored in a pointer, but the pointer itself (which holds the address) is stored in the stack.
- The address (stored in the pointer) points to the location of the value, and you cannot change the address itself to point to something else after it's assigned (no pointer arithmetic in Go).

Pointer Arithmetic in C vs Go:
- In C, pointer arithmetic is allowed, which means you can directly modify the memory address stored in a pointer, which can be dangerous if not done carefully.
- Go does not allow pointer arithmetic. A pointer in Go is simply an address of a variable and cannot be used to navigate through memory (e.g., moving the pointer from one location to another).

Stack and Heap:
- In Go, the Go runtime decides whether to allocate memory on the stack or heap. Typically:
  - Small variables, like `int` and `bool`, are allocated on the stack.
  - Larger or more complex variables, such as slices, maps, and dynamically sized structs, are allocated on the heap.

  Example:
    var s []int = []int{1, 2, 3}  // 's' is a slice, dynamically sized, so it is allocated on the heap.

  - You cannot directly control whether a variable goes to the stack or heap, but you can control what is allocated by understanding Go's memory management model.

Stack Memory:
- The stack grows and shrinks as functions are called and returned.
- Each function call creates a new stack frame, where local variables are stored.
- Once the function returns, its stack frame is popped, and the memory is reclaimed.

Heap Memory:
- Memory allocated in the heap remains allocated until it is explicitly garbage collected.
- Go uses a garbage collector to automatically reclaim unused memory (heap) to avoid memory leaks.
*/
