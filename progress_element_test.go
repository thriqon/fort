package fort

import "fmt"

//current, total uint64, width int) string {

func ExampleProgressAsBar() {
	fmt.Println(progressAsBar(0, 100, 10))
	fmt.Println(progressAsBar(100, 100, 10))
	fmt.Println(progressAsBar(80, 100, 10))
	// Output:
	// |>-------|
	// |========|
	// |=====>--|
}
