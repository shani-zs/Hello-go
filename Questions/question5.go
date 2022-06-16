package Questions

import "fmt"

func Question5() {
	if v := 5 * 10; v > 20 {
		fmt.Println("v>20")
	} else {
		fmt.Println("v<20")
	}
	fmt.Println(v)
}
