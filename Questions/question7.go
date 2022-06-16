package Questions

import "fmt"

func Question7() {
	var i interface{} = 10

	s := i.(string)
	fmt.Println(s)

}
