package composites

import "fmt"

// Base struct
type Entity struct {
	ID string
}

func (e Entity) GetID() string {
	return e.ID
}

// Employee embeds Entity. This means an Employ "has an" Entity.
// This fields and methods of Entity are "Promoted" to Employee.
type Employee struct {
	Entity // Embedding Entity struct
	Name   string
	Role   string
}

// DemonstrateEmbedding shows composition via struct Embedding.
func DemonstrateEmbedding() {
	fmt.Println("\n[Struct Embedding (Composition)]")

	emp := Employee{
		Entity: Entity{ID: "123"},
		Name:   "Bob",
		Role:   "Engineer",
	}

	// We can access the embedded Entity's fields and methods directly.
	fmt.Printf("Employee Name: %s, Role: %s\n", emp, emp.Name, emp.Role)

	// We can also call the embedded struct's methods directly.
	fmt.Printf("Employee ID: %s\n", emp.GetID())
}
