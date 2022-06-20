package Questions

import "fmt"

func mul(i, j int) int {

	defer fmt.Println("func mul is called")
	res := i * j
	fmt.Println("result:", res)
	return 0
}

func show() {

	fmt.Println("Hello!, Zopsamrt")
}

func Question20() {
	mul(23, 45)

	defer mul(23, 56)

	defer show()
}

func Question23() {

	var s *int
	fmt.Println("s=", s)

	i := 2.45
	j := &i
	fmt.Printf("%v, %T", *j, j)

}

func Question24() {
	i := 42
	fmt.Printf("i: %[1]T %[1]d\n", i)
	p := &i
	fmt.Printf("p: %[1]T %[1]p\n", p)

	j := *p
	fmt.Printf("j: %[1]T %[1]d\n", j)

	q := &p
	fmt.Printf("q: %[1]T %[1]p\n", q)

	k := **q
	fmt.Printf("k: %[1]T %[1]d\n", k)
}
