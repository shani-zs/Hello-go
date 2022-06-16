package Questions

import "fmt"

func Question6() {
	v := []int{1, 2, 3, 4}

	if v == nil {
		fmt.Println("nil")
	}
	if v == []int{1, 2, 3, 4} {
		fmt.Println("equal")
	}
}
