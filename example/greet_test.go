package example

import "fmt"

func ExampleHello() {
	greeting, err := Hello("Jon")
	if err != nil {
		panic(err)
	}
	fmt.Println(greeting)

	// Output:
	// Hello, Jon
}

func ExamplePage() {
	checkIns := map[string]bool{
		"Bob":    true,
		"Jack":   false,
		"Tony":   false,
		"Maggie": true,
		"Simon":  false,
		"Becky":  false,
		"Jon":    true,
	}
	Page(checkIns)

	// Unordered Output:
	// Paging Jack; please see the front desk to check in.
	// Paging Tony; please see the front desk to check in.
	// Paging Simon; please see the front desk to check in.
	// Paging Becky; please see the front desk to check in.
}
