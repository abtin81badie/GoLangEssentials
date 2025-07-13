package advance

import "fmt"

// iota is a constant generator that increments on each line within a const block.
const (
	Sunday  = iota // 0
	Monday         // 1
	Tuesday        // 2
)

// iota can be used in expressions.
const (
	_  = 1 << (10 * iota) // Ignore first value
	KB                    // 1 << (10*1)
	MB                    // 1 << (10*2)
	GB                    // 1 << (10*3)
)

// DemonstrateIota shows different uses of the iota keyword.
func DemonstrateIota() {
	fmt.Println("\n[Iota]")
	fmt.Printf("Sunday: %d, Monday: %d, Tuesday: %d\n", Sunday, Monday, Tuesday)
	fmt.Printf("KB: %d, MB: %d, GB: %d\n", KB, MB, GB)
}
