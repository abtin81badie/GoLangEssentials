package main

import (
	"fmt"

	"github.com/abtin81badie/GoLangEssentials/go-cheatsheet/advance"
	"github.com/abtin81badie/GoLangEssentials/go-cheatsheet/basics"
	"github.com/abtin81badie/GoLangEssentials/go-cheatsheet/collections"
	"github.com/abtin81badie/GoLangEssentials/go-cheatsheet/composites"
	"github.com/abtin81badie/GoLangEssentials/go-cheatsheet/concurrency"
	"github.com/abtin81badie/GoLangEssentials/go-cheatsheet/generics"
)

func main() {
	// This is the main entry point of the Go application.
	// You can add your code here to implement the functionality you need.
	fmt.Println("Hello, Go Cheatsheet!")

	// -- Basics --
	fmt.Println("--- 1. Basics: Variables, Control Flow, Functions ---")
	basics.DemonstrateVariables()
	basics.DemonstrateControlFlow()
	basics.DemonstrateFunctions()

	// -- Collections --
	fmt.Println("--- 2. Collections: Arrays, Slices, Maps ---")
	collections.DemonstrateArraysAndSlices()
	collections.DemonstrateMaps()

	// -- COMPOSITE TYPES --
	fmt.Println("--- 3. Composite Types: Structs, Methods ,Interfaces ---")
	composites.DemonstrateStructsMethods()
	composites.DemonstrateInterfaces()
	composites.DemonstrateEmbedding()

	// --- ADVANCED CONCEPTS ---
	fmt.Println("\n--- 4. Advanced Concepts ---")
	advance.DemonstrateClosures()
	advance.DemonstrateErrors()
	advance.DemonstratePointers()
	advance.DemonstrateIota()
	advance.DemonstrateDefer()

	// --- CONCURRENCY ---
	fmt.Println("\n--- 5. Concurrency ---")
	concurrency.DemonstrateGoroutinesAndChannels()

	// --- GENERICS ---
	fmt.Println("\n--- 6. Generics ---")
	generics.DemonstrateGenerics()
}
