// main.go
package main

import (
	"fmt"
	"go-concurrency-project/advancedchannels"
	"go-concurrency-project/atomics"
	"go-concurrency-project/barriers"
	"go-concurrency-project/channels"
	"go-concurrency-project/condvars"
	"go-concurrency-project/goroutines"
	"go-concurrency-project/mutexes"
	"go-concurrency-project/raceconditions"
	"go-concurrency-project/selects"
	"go-concurrency-project/waitgroups"
)

func main() {
	fmt.Println("--- Go Concurrency Tutorial ---")
	fmt.Println("Uncomment the example you want to run in main.go")

	// To run an example, uncomment its corresponding line.
	// It's recommended to run them one at a time to see the output clearly.

	goroutines.RunExample()
	waitgroups.RunExample()
	raceconditions.RunExample() // Remember to run this with the -race flag!
	mutexes.RunExample()
	atomics.RunExample()
	condvars.RunExample()
	barriers.RunExample()
	channels.RunExample()
	selects.RunExample()
	advancedchannels.RunExample()
}
