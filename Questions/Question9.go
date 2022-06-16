package Questions

import "fmt"

func Question9() {
	i := 0

	defer fmt.Println("start")
	defer fmt.Printf("defer 1 i: %d\n", i)

	i++

	defer fmt.Printf("defer 2 i: %d\n", i)
	defer fmt.Printf("end\n")
}
