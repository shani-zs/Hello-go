package Questions

import "fmt"

func Question8() {
	v := [4]int{1, 2, 3, 4}

	switch v {
	case [4]int{}:
		fmt.Println("case1")
	default:
		fmt.Println("default case")
	case [4]int{1, 2, 3, 4}:
		fmt.Println("case 2")
	}
}

//printing case 2
