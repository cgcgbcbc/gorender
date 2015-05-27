package main

import (
	"os"
)

func Example_string() {
	os.Args = []string{"gorender", "--string", "{{.Count}}", "Count=1"}
	main()
	// Output:
	// 1
}

func Example_file() {
	os.Args = []string{"gorender", "--path", "./fixtures/tmpl.txt", "Count=1"}
	main()
	// Output:
	// 1
}
