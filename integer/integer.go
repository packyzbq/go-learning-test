package integer

import "fmt"

// Add is a method for adding
func Add(a int, b int) (r int) {
	return a + b
}

func main() {
	fmt.Println(Add(1, 2))
}
