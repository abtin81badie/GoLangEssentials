package main

import (
	"fmt"
)

// Global Variables Declaration
var (
	Age    int    // Variable to store age
	Name   string // Variable to store name
	Family string // Variable to store family name
)

// Constants Declaration
const (
	MAX_COUNT  = 1000          // Constant for maximum count
	NICK_NAME  = "Erren Yager" // Constant for nickname
	LET_PROCED = false         // Constant boolean value
)

func main() {
	// Greet the user
	fmt.Println("Hello, World!")

	// Assign a value to 'Name' variable and demonstrate pass-by-value and pass-by-reference
	Name := "Abtin" // Local variable 'Name' shadows the global 'Name' variable

	// Pass-by-value: Original variable 'Name' is unchanged
	passByValue(Name)
	fmt.Println("After passByValue, Name:", Name)

	// Pass-by-reference: The original variable 'Name' is modified
	passByReference(&Name)
	fmt.Println("After passByReference, Name:", Name)

	// Demonstrate control structures
	demonstrateControlStructures()

	// Function to demonstrate different variable types and their initialization
	variables()

	// Demonstrating constant usage
	const TEST string = "A"
	fmt.Println("The constant TEST is:", TEST)

	// Arrays: Fixed-size, can be accessed with an index
	var a1 [10]int  // Declare an array with default values
	a2 := [10]int{} // Declare and initialize an array with zero values
	fmt.Println("a1[7]:", a1[7])
	fmt.Println("a2[0]:", a2[0])

	// Slices: Dynamically-sized, can be resized with append
	var s1 []int // Declare an empty slice (nil slice)
	// fmt.Println(s1[90]) // Uncommenting this will panic, as the slice is empty
	s1 = append(s1, 10)   // Append a value to the slice
	s2 := make([]int, 10) // Create a slice with a specified length
	s2 = append(s2, 11)   // Append a value to the slice
	fmt.Println("s1[0]:", s1[0])
	fmt.Println("s1 length, capacity:", len(s1), cap(s1))
	fmt.Println("s2[10]:", s2[10])
	fmt.Println("s2 length, capacity:", len(s2), cap(s2))

	// Maps: Unordered collection of key-value pairs
	// Keys must be comparable (string, int, etc.)
	m1 := make(map[string]string) // Initialize an empty map
	m2 := map[string]int{         // Initialize a map with values
		"Abtin":   22,
		"Fatemeh": 22,
	}

	var m3 map[int]bool // Uninitialized map (nil)
	fmt.Println("m2[\"Abtin\"]:", m2["Abtin"])
	fmt.Println("m2[\"Fatemeh\"]:", m2["Fatemeh"])

	// Demonstrating map assignments
	m1["hello"] = "world"
	fmt.Printf("m1: %v\n", m1)

	// Demonstrating a nil map
	fmt.Printf("m3: %v\n", m3)

	// Demonstrating map item existence and deletion
	m := map[string]string{
		"Name": "Abtin",
	}

	_, ok := m["Name"]
	if ok {
		println("Item exists")
	} else {
		println("Item does not exist")
	}

	val, ok := m["Name"]
	if ok {
		println("Item exists.")
		println(val)
	} else {
		println("Item does not exist.")
	}

	delete(m, "Name")

	_, ok = m["Name"]
	if ok {
		println("Item exists")
	} else {
		println("Item does not exist")
	}

	// Struct and Methods Example
	person := Person{Name: "Abtin", Age: 22}
	fmt.Println(person.Introduce())

	p1 := Person{"Fatemeh", 21}
	fmt.Println(p1.Introduce())

	var p2 Person
	p2.Name = "Amir"
	p2.Age = 21
	fmt.Println(p2.Introduce())

	// Initialize by Zero Values
	p3 := new(Person)
	fmt.Println(p3.Introduce())

	// nil Initialized
	// var p4 *Person
	// fmt.Println(p4.Introduce())

	// Pointers
	pointerExample()

	// Interface Example
	demoInterfaceExample()

	// Concurrency Example
	go demoConcurrency()

	// Error Handling Example
	if err := demoErrorHandling(4); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Number is even")
	}

	// Defer Example
	demoDefer()

	// Type Conversion Example
	demoTypeConversion()

	// Shift operation
	leftShift()
	rightShift()
}

// Function that demonstrates pass-by-value (does not modify the original value)
func passByValue(s string) {
	s = "felan" // Modifies only the local variable
}

// Function that demonstrates pass-by-reference (modifies the original value)
func passByReference(s *string) {
	*s = "felan" // Modifies the original variable through the pointer
}

// Function to demonstrate different variable declarations and types
func variables() {
	// Declare a variable with a specific type
	var var1 int32 = 32
	fmt.Println("var1 (int32):", var1)

	// Short declaration (type inferred)
	var2 := "string"
	fmt.Println("var2 (string):", var2)

	// Short declaration for a float
	var3 := 32.2
	fmt.Println("var3 (float64):", var3)

	// Multiple variable declaration
	v1, v2, v3 := false, "hello", 23
	fmt.Println("v1 (bool), v2 (string), v3 (int):", v1, v2, v3)
}

// Structs and Methods Example
type Person struct {
	Name string
	Age  int
}

func (p Person) Introduce() string {
	return fmt.Sprintf("Hi, I'm %s, %d years old.", p.Name, p.Age)
}

// Interface Example
type Speaker interface {
	Introduce() string
}

func demoInterfaceExample() {
	var speaker Speaker = Person{Name: "Abtin", Age: 22}
	fmt.Println(speaker.Introduce())
}

// Goroutines and Channels Example
func demoConcurrency() {
	ch := make(chan string)
	go func() {
		ch <- "Hello from Goroutine!"
	}()
	message := <-ch
	fmt.Println(message)
}

// Error Handling Example
func demoErrorHandling(n int) error {
	if n%2 != 0 {
		return fmt.Errorf("number is odd")
	}
	return nil
}

// Defer Example
func demoDefer() {
	defer fmt.Println("This will be printed last!")
	fmt.Println("This will be printed first!")
}

// Type Conversion Example
func demoTypeConversion() {
	var intVar int = 42
	var floatVar float64 = float64(intVar)
	fmt.Println("Converted int to float:", floatVar)

	var strVar string = fmt.Sprintf("%d", intVar)
	fmt.Println("Converted int to string:", strVar)
}

// Control Structures Example
func demonstrateControlStructures() {
	// If-else statement with initialization
	if x := 10; x > 5 {
		fmt.Println("x is greater than 5")
	} else {
		fmt.Println("x is less than or equal to 5")
	}

	// Switch statement
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

	// For loops
	for i := 0; i < 5; i++ {
		fmt.Println("For loop iteration:", i)
	}

	// Infinite loop
	// Uncomment to test
	// for {
	//     fmt.Println("This loop runs forever")
	// }
}

func leftShift() {
	// Left Shift Example
	var a int = 4       // In binary: 0100
	var result = a << 2 // Shifting left by 2 positions

	// Binary representation before and after the shift
	fmt.Printf("a: %b(binary), %d (decimal)\n", a, a)
	fmt.Printf("a << 2: %b(binary), %d (decimal)\n", result, result) // Expected: 16 (binary 10000)
}

func rightShift() {
	// Right Shift Example
	var a int = 16      // In binary: 10000
	var result = a >> 2 // Shifting right by 2 positions

	// Binary representation before and after the shift
	fmt.Printf("a: %b (binary), %d (decimal)\n", a, a)
	fmt.Printf("a >> 2: %b (binary), %d (decimal)\n", result, result)
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

Example:
*/

func pointerExample() {
	// 'a' is a primitive variable stored in the stack
	var a int = 100
	var ptr *int // 'ptr' is a pointer to an integer
	ptr = &a     // 'ptr' now holds the address of 'a'

	// Dereferencing ptr to access the value at the address
	fmt.Println("Address of 'a':", ptr)            // Prints the values of 'a'
	fmt.Println("value of 'a' through ptr:", *ptr) // Prints the value of 100

	// Create a pointer to a Person struct and change its Name using the pointer
	p := &Person{Name: "Parsa", Age: 21}
	changeName(p)
	fmt.Println("Updated Name:", p.Name) // Prints the updated name "Batman"

	// Example of safe usage of pointers and dereferencing
	var x *uint // x is a nil pointer
	var y uint  // y is a regular uint variable

	// Initialization y to a value before dereferencing
	y = 10
	x = &y // Now x points to y

	// Dereferencing x to get the value it points to
	fmt.Println("x points to value:", *x) // Prints the value 10

	// Demonstrating pointers to pointers
	z := &y                                   // z is a pointer wto y
	fmt.Println("Address of y through z:", z) // Prints the address of y
	fmt.Println("Value of y through z:", *z)  // Prints the value of 10
}

func changeName(p *Person) {
	// Modify the Name field of Person struct through the pointer
	p.Name = "Batman"
}

/*
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
