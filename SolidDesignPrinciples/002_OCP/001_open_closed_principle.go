package main

import "fmt"

type Color int

const (
	red Color = iota
	green
	blue
)

type Size int

const (
	small Size = iota
	medium
	large
)

type Product struct {
	name  string
	color Color
	size  Size
}

/*
	Not following Proper OCP
	Each time requirement changes we are adding New Method and duplication of Code
	Extension is not possible
*/
type Filter struct{}

// First Requirement Filter By Color
func (f *Filter) filterByColor(products []Product, color Color) []*Product {
	result := make([]*Product, 0)

	for i, v := range products {
		if v.color == color {
			result = append(result, &products[i])
		}
	}
	return result
}

// Second Requirement Filter By Size
func (f *Filter) filterBySize(products []Product, size Size) []*Product {
	result := make([]*Product, 0)

	for i, v := range products {
		if v.size == size {
			result = append(result, &products[i])
		}
	}
	return result
}

// Second Requirement Filter By Size
func (f *Filter) filterByColorAndSize(products []Product, color Color, size Size) []*Product {
	result := make([]*Product, 0)

	for i, v := range products {
		if v.color == color && v.size == size {
			result = append(result, &products[i])
		}
	}
	return result
}

/*
	Following OCP and Specification Pattern
	Easily extend the functionality and code duplication removed
	Can be achieved through Interface
*/

type Specification interface {
	IsSatisfied(p *Product) bool
}

// First requirement is Filter By Color
type ColorSpecification struct {
	color Color
}

func (spec ColorSpecification) IsSatisfied(p *Product) bool {
	return p.color == spec.color
}

// Second requirement is Filter By Size
type SizeSpecification struct {
	size Size
}

func (spec SizeSpecification) IsSatisfied(p *Product) bool {
	return p.size == spec.size
}

// Third requirement is Filter By Color and Size
type ColorSizeSpecification struct {
	color, size Specification
}

func (spec ColorSizeSpecification) IsSatisfied(p *Product) bool {
	return spec.color.IsSatisfied(p) && spec.size.IsSatisfied(p)
}

type BetterFilter struct {}

func (f *BetterFilter) Filter(products []Product, spec Specification) []*Product {
	result := make([]*Product, 0)
	for i, v := range products {
		if spec.IsSatisfied(&v) {
			result = append(result, &products[i])
		}
	}
	return result
} 

/*
	Following OCP
	Embedding is a powerful tool which allows Go's types to be open for extension.
	Two Examples Given below.
*/

// Example 1

// Requirement 1 - Display Name (Requirement Delivered to 100 Customers i.e. Running In Production)
type Employee struct {
	name string
}

func (e Employee) Display() {
	fmt.Printf("Hello %s\n", e.name)
}

// Requirement 2 - Display Name with SurName (As Requirement 1 already in Production 
// so don't disturb earlier implmentation)
// This Step called as 
// Open For Extension but closed for modification
type EmployeeWithSurname struct {
	Employee
	surname string
}

func (e EmployeeWithSurname) Display() {
	fmt.Printf("Hello %s %s\n", e.name, e.surname)
}


func main() {
	item001 := Product{"Item 001", green, small}
	item002 := Product{"Item 002", green, large}
	item003 := Product{"Item 003", blue, small}
	item004 := Product{"Item 004", green, small}

	products := []Product{item001, item002, item003, item004}

	fmt.Print("Green products (old): \n")
	f := Filter{}
	// Filter by Color
	for _, v := range f.filterByColor(products, green) {
		fmt.Printf(" - %s is green\n", v.name)
	}

	// Following Proper OCP and Specification Pattern
	fmt.Print("Green products (new):\n")
	greenSpec := ColorSpecification{green}
	betterf := BetterFilter{}
	for _, v := range betterf.Filter(products, greenSpec) {
		fmt.Printf(" - %s is green\n", v.name)
	}

	fmt.Print("Large products (new):\n")
	largeSpec := SizeSpecification{large}
	for _, v := range betterf.Filter(products, largeSpec) {
		fmt.Printf(" - %s is large\n", v.name)
	}
	
	fmt.Print("Large blue items:\n")
	largeGreenSpec := ColorSizeSpecification{largeSpec, greenSpec}
	for _, v := range betterf.Filter(products, largeGreenSpec) {
		fmt.Printf(" - %s is large and green\n", v.name)
	}

	// Example 1

	// Requirement 1 - Display Name (Requirement Delivered to 100 Customers i.e. Running In Production)
	displayName := Employee{"Name"}
	displayName.Display()

	// Requirement 2 - Display Name with SurName (As Requirement 1 already in Production 
	// so don't disturb earlier implmentation)
	// This Step called as 
	// Open For Extension but closed for modification
	displayNameWithSurname := EmployeeWithSurname{displayName, "Surname"}
	displayNameWithSurname.Display()
}
